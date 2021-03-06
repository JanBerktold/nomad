package e2eutil

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/helper"
	"github.com/hashicorp/nomad/jobspec"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/hashicorp/nomad/testutil"
	"github.com/kr/pretty"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/require"
)

// retries is used to control how many times to retry checking if the cluster has a leader yet
const retries = 500

func WaitForLeader(t *testing.T, nomadClient *api.Client) {
	statusAPI := nomadClient.Status()

	testutil.WaitForResultRetries(retries, func() (bool, error) {
		leader, err := statusAPI.Leader()
		return leader != "", err
	}, func(err error) {
		t.Fatalf("failed to find leader: %v", err)
	})
}

// WaitForNodesReady waits until at least `nodes` number of nodes are ready or
// fails the test.
func WaitForNodesReady(t *testing.T, nomadClient *api.Client, nodes int) {
	nodesAPI := nomadClient.Nodes()

	testutil.WaitForResultRetries(retries, func() (bool, error) {
		defer time.Sleep(time.Millisecond * 100)
		nodesList, _, err := nodesAPI.List(nil)
		if err != nil {
			return false, fmt.Errorf("error listing nodes: %v", err)
		}

		eligibleNodes := 0
		for _, node := range nodesList {
			if node.Status == "ready" {
				eligibleNodes++
			}
		}

		return eligibleNodes >= nodes, fmt.Errorf("only %d nodes ready (wanted at least %d)", eligibleNodes, nodes)
	}, func(err error) {
		t.Fatalf("failed to get enough ready nodes: %v", err)
	})
}

func stringToPtrOrNil(s string) *string {
	if s == "" {
		return nil
	}
	return helper.StringToPtr(s)
}

func RegisterAllocs(t *testing.T, nomadClient *api.Client, jobFile, jobID, cToken string) []*api.AllocationListStub {
	r := require.New(t)

	// Parse job
	job, err := jobspec.ParseFile(jobFile)
	r.Nil(err)

	// Set custom job ID (distinguish among tests)
	job.ID = helper.StringToPtr(jobID)

	// Set a Consul "operator" token for the job, if provided.
	job.ConsulToken = stringToPtrOrNil(cToken)

	// Register job
	var idx uint64
	jobs := nomadClient.Jobs()
	testutil.WaitForResult(func() (bool, error) {
		resp, meta, err := jobs.Register(job, nil)
		if err != nil {
			return false, err
		}
		idx = meta.LastIndex
		return resp.EvalID != "", fmt.Errorf("expected EvalID:%s", pretty.Sprint(resp))
	}, func(err error) {
		r.NoError(err)
	})

	allocs, _, err := jobs.Allocations(jobID, false, &api.QueryOptions{WaitIndex: idx})
	require.NoError(t, err)
	return allocs
}

func RegisterAndWaitForAllocs(t *testing.T, nomadClient *api.Client, jobFile, jobID, cToken string) []*api.AllocationListStub {
	r := require.New(t)
	g := NewGomegaWithT(t)
	jobs := nomadClient.Jobs()

	// Start allocations
	RegisterAllocs(t, nomadClient, jobFile, jobID, cToken)

	// Wrap in retry to wait until placement
	g.Eventually(func() []*api.AllocationListStub {
		// Look for allocations
		allocs, _, _ := jobs.Allocations(jobID, false, nil)
		return allocs
	}, 30*time.Second, time.Second).ShouldNot(BeEmpty())

	allocs, _, err := jobs.Allocations(jobID, false, nil)
	r.NoError(err)
	return allocs
}

func WaitForAllocRunning(t *testing.T, nomadClient *api.Client, allocID string) {
	testutil.WaitForResultRetries(retries, func() (bool, error) {
		time.Sleep(time.Millisecond * 100)
		alloc, _, err := nomadClient.Allocations().Info(allocID, nil)
		if err != nil {
			return false, err
		}

		return alloc.ClientStatus == structs.AllocClientStatusRunning, fmt.Errorf("expected status running, but was: %s", alloc.ClientStatus)
	}, func(err error) {
		t.Fatalf("failed to wait on alloc: %v", err)
	})
}

func WaitForAllocsRunning(t *testing.T, nomadClient *api.Client, allocIDs []string) {
	for _, allocID := range allocIDs {
		WaitForAllocRunning(t, nomadClient, allocID)
	}
}

func WaitForAllocsNotPending(t *testing.T, nomadClient *api.Client, allocIDs []string) {
	for _, allocID := range allocIDs {
		WaitForAllocNotPending(t, nomadClient, allocID)
	}
}

func WaitForAllocNotPending(t *testing.T, nomadClient *api.Client, allocID string) {
	testutil.WaitForResultRetries(retries, func() (bool, error) {
		time.Sleep(time.Millisecond * 100)
		alloc, _, err := nomadClient.Allocations().Info(allocID, nil)
		if err != nil {
			return false, err
		}

		return alloc.ClientStatus != structs.AllocClientStatusPending, fmt.Errorf("expected status not pending, but was: %s", alloc.ClientStatus)
	}, func(err error) {
		t.Fatalf("failed to wait on alloc: %v", err)
	})
}

func AllocIDsFromAllocationListStubs(allocs []*api.AllocationListStub) []string {
	allocIDs := make([]string, 0, len(allocs))
	for _, alloc := range allocs {
		allocIDs = append(allocIDs, alloc.ID)
	}
	return allocIDs
}

func DeploymentsForJob(t *testing.T, nomadClient *api.Client, jobID string) []*api.Deployment {
	ds, _, err := nomadClient.Deployments().List(nil)
	require.NoError(t, err)

	out := []*api.Deployment{}
	for _, d := range ds {
		if d.JobID == jobID {
			out = append(out, d)
		}
	}

	return out
}

func WaitForDeployment(t *testing.T, nomadClient *api.Client, deployID string, status string, statusDesc string) {
	testutil.WaitForResultRetries(retries, func() (bool, error) {
		time.Sleep(time.Millisecond * 100)
		deploy, _, err := nomadClient.Deployments().Info(deployID, nil)
		if err != nil {
			return false, err
		}

		if deploy.Status == status && deploy.StatusDescription == statusDesc {
			return true, nil
		}
		return false, fmt.Errorf("expected status %s \"%s\", but got: %s \"%s\"",
			status,
			statusDesc,
			deploy.Status,
			deploy.StatusDescription,
		)

	}, func(err error) {
		t.Fatalf("failed to wait on deployment: %v", err)
	})
}

// fake is a package that includes fake implementations of public interfaces
// from the CSI package for testing.
package fake

import (
	"context"
	"errors"
	"sync"

	"github.com/hashicorp/nomad/plugins/base"
	"github.com/hashicorp/nomad/plugins/csi"
	"github.com/hashicorp/nomad/plugins/shared/hclspec"
)

var _ csi.CSIPlugin = &Client{}

// Client is a mock implementation of the csi.CSIPlugin interface for use in testing
// external components
type Client struct {
	Mu sync.RWMutex

	NextPluginInfoResponse *base.PluginInfoResponse
	NextPluginInfoErr      error
	PluginInfoCallCount    int64

	NextPluginProbeResponse bool
	NextPluginProbeErr      error
	PluginProbeCallCount    int64

	NextPluginGetInfoResponse string
	NextPluginGetInfoErr      error
	PluginGetInfoCallCount    int64

	NextPluginGetCapabilitiesResponse *csi.PluginCapabilitySet
	NextPluginGetCapabilitiesErr      error
	PluginGetCapabilitiesCallCount    int64

	NextNodeGetInfoResponse *csi.NodeGetInfoResponse
	NextNodeGetInfoErr      error
	NodeGetInfoCallCount    int64
}

// PluginInfo describes the type and version of a plugin.
func (c *Client) PluginInfo() (*base.PluginInfoResponse, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.PluginInfoCallCount++

	return c.NextPluginInfoResponse, c.NextPluginInfoErr
}

// ConfigSchema returns the schema for parsing the plugins configuration.
func (c *Client) ConfigSchema() (*hclspec.Spec, error) {
	return nil, errors.New("Unsupported")
}

// SetConfig is used to set the configuration by passing a MessagePack
// encoding of it.
func (c *Client) SetConfig(a *base.Config) error {
	return errors.New("Unsupported")
}

// PluginProbe is used to verify that the plugin is in a healthy state
func (c *Client) PluginProbe(ctx context.Context) (bool, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.PluginProbeCallCount++

	return c.NextPluginProbeResponse, c.NextPluginProbeErr
}

// PluginGetInfo is used to return semantic data about the plugin.
// Response:
//  - string: name, the name of the plugin in domain notation format.
func (c *Client) PluginGetInfo(ctx context.Context) (string, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.PluginGetInfoCallCount++

	return c.NextPluginGetInfoResponse, c.NextPluginGetInfoErr
}

// PluginGetCapabilities is used to return the available capabilities from the
// identity service. This currently only looks for the CONTROLLER_SERVICE and
// Accessible Topology Support
func (c *Client) PluginGetCapabilities(ctx context.Context) (*csi.PluginCapabilitySet, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.PluginGetCapabilitiesCallCount++

	return c.NextPluginGetCapabilitiesResponse, c.NextPluginGetCapabilitiesErr
}

// NodeGetInfo is used to return semantic data about the current node in
// respect to the SP.
func (c *Client) NodeGetInfo(ctx context.Context) (*csi.NodeGetInfoResponse, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.NodeGetInfoCallCount++

	return c.NextNodeGetInfoResponse, c.NextNodeGetInfoErr
}

// Shutdown the client and ensure any connections are cleaned up.
func (c *Client) Close() error {
	return nil
}
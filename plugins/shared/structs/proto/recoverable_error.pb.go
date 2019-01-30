// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/shared/structs/proto/recoverable_error.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RecoverableError is used with a grpc Status to indicate if the error is one
// which is recoverable and can be reattempted by the client.
type RecoverableError struct {
	Recoverable          bool     `protobuf:"varint,1,opt,name=recoverable,proto3" json:"recoverable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoverableError) Reset()         { *m = RecoverableError{} }
func (m *RecoverableError) String() string { return proto.CompactTextString(m) }
func (*RecoverableError) ProtoMessage()    {}
func (*RecoverableError) Descriptor() ([]byte, []int) {
	return fileDescriptor_recoverable_error_8c5d7f86073ca00c, []int{0}
}
func (m *RecoverableError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecoverableError.Unmarshal(m, b)
}
func (m *RecoverableError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecoverableError.Marshal(b, m, deterministic)
}
func (dst *RecoverableError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoverableError.Merge(dst, src)
}
func (m *RecoverableError) XXX_Size() int {
	return xxx_messageInfo_RecoverableError.Size(m)
}
func (m *RecoverableError) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoverableError.DiscardUnknown(m)
}

var xxx_messageInfo_RecoverableError proto.InternalMessageInfo

func (m *RecoverableError) GetRecoverable() bool {
	if m != nil {
		return m.Recoverable
	}
	return false
}

func init() {
	proto.RegisterType((*RecoverableError)(nil), "hashicorp.nomad.plugins.shared.structs.RecoverableError")
}

func init() {
	proto.RegisterFile("plugins/shared/structs/proto/recoverable_error.proto", fileDescriptor_recoverable_error_8c5d7f86073ca00c)
}

var fileDescriptor_recoverable_error_8c5d7f86073ca00c = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0x2b, 0xd6, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0xd1, 0x2f, 0x2e, 0x29, 0x2a, 0x4d,
	0x2e, 0x29, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4b, 0x2d,
	0x4a, 0x4c, 0xca, 0x49, 0x8d, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x03, 0x8b, 0x0b, 0xa9, 0x65,
	0x24, 0x16, 0x67, 0x64, 0x26, 0xe7, 0x17, 0x15, 0xe8, 0xe5, 0xe5, 0xe7, 0x26, 0xa6, 0xe8, 0x41,
	0x4d, 0xd1, 0x83, 0x98, 0xa2, 0x07, 0x35, 0x45, 0xc9, 0x84, 0x4b, 0x20, 0x08, 0x61, 0x84, 0x2b,
	0xc8, 0x04, 0x21, 0x05, 0x2e, 0x6e, 0x24, 0x63, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x90,
	0x85, 0x9c, 0xd8, 0xa3, 0x58, 0xc1, 0xd6, 0x24, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0x45, 0x79, 0xed, 0xa5, 0x00, 0x00, 0x00,
}

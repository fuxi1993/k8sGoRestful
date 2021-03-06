// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pods_requests.proto

package models

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PodsResponse struct {
	ShouldStart      bool  `protobuf:"varint,1,opt,name=should_start,json=shouldStart" json:"should_start,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PodsResponse) Reset()                    { *m = PodsResponse{} }
func (m *PodsResponse) String() string            { return proto.CompactTextString(m) }
func (*PodsResponse) ProtoMessage()               {}
func (*PodsResponse) Descriptor() ([]byte, []int) { return fileDescriptorPodsRequests, []int{0} }

func (m *PodsResponse) GetShouldStart() bool {
	if m != nil {
		return m.ShouldStart
	}
	return false
}

func init() {
	proto.RegisterType((*PodsResponse)(nil), "models.PodsResponse")
}

func init() { proto.RegisterFile("pods_requests.proto", fileDescriptorPodsRequests) }

var fileDescriptorPodsRequests = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc8, 0x4f, 0x29,
	0x8e, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0xcb, 0xcd, 0x4f, 0x49, 0xcd, 0x29, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79,
	0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa6, 0x64, 0xc8, 0xc5, 0x13, 0x90, 0x9f, 0x52, 0x1c, 0x94, 0x5a,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xc8, 0xc5, 0x53, 0x9c, 0x91, 0x5f, 0x9a, 0x93, 0x12,
	0x5f, 0x5c, 0x92, 0x58, 0x54, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x0d, 0x11, 0x0b,
	0x06, 0x09, 0x01, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0xf0, 0x1a, 0x93, 0x7f, 0x00, 0x00, 0x00,
}

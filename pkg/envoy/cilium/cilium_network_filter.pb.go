// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/cilium_network_filter.proto

package cilium

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkFilter struct {
	// Path to the proxylib to be opened
	Proxylib string `protobuf:"bytes,1,opt,name=proxylib,proto3" json:"proxylib,omitempty"`
	// Transparent set of parameters provided for proxylib initialization
	ProxylibParams map[string]string `protobuf:"bytes,2,rep,name=proxylib_params,json=proxylibParams,proto3" json:"proxylib_params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// L7 protocol identifier
	L7Proto string `protobuf:"bytes,3,opt,name=l7_proto,json=l7Proto,proto3" json:"l7_proto,omitempty"`
	// Name of the policy to be enforced
	PolicyName           string   `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkFilter) Reset()         { *m = NetworkFilter{} }
func (m *NetworkFilter) String() string { return proto.CompactTextString(m) }
func (*NetworkFilter) ProtoMessage()    {}
func (*NetworkFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_97d495a4248b4ff1, []int{0}
}

func (m *NetworkFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkFilter.Unmarshal(m, b)
}
func (m *NetworkFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkFilter.Marshal(b, m, deterministic)
}
func (m *NetworkFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkFilter.Merge(m, src)
}
func (m *NetworkFilter) XXX_Size() int {
	return xxx_messageInfo_NetworkFilter.Size(m)
}
func (m *NetworkFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkFilter.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkFilter proto.InternalMessageInfo

func (m *NetworkFilter) GetProxylib() string {
	if m != nil {
		return m.Proxylib
	}
	return ""
}

func (m *NetworkFilter) GetProxylibParams() map[string]string {
	if m != nil {
		return m.ProxylibParams
	}
	return nil
}

func (m *NetworkFilter) GetL7Proto() string {
	if m != nil {
		return m.L7Proto
	}
	return ""
}

func (m *NetworkFilter) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkFilter)(nil), "cilium.NetworkFilter")
	proto.RegisterMapType((map[string]string)(nil), "cilium.NetworkFilter.ProxylibParamsEntry")
}

func init() { proto.RegisterFile("cilium/cilium_network_filter.proto", fileDescriptor_97d495a4248b4ff1) }

var fileDescriptor_97d495a4248b4ff1 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xce, 0xcc, 0xc9,
	0x2c, 0xcd, 0xd5, 0x87, 0x50, 0xf1, 0x79, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0xf1, 0x69, 0x99,
	0x39, 0x25, 0xa9, 0x45, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x49, 0xa5, 0x7f,
	0x8c, 0x5c, 0xbc, 0x7e, 0x10, 0x05, 0x6e, 0x60, 0x79, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xa2, 0xfc,
	0x8a, 0xca, 0x9c, 0xcc, 0x24, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x5f, 0x28, 0x88,
	0x8b, 0x1f, 0xc6, 0x8e, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0x96, 0x60, 0x52, 0x60, 0xd6, 0xe0,
	0x36, 0xd2, 0xd4, 0x83, 0x98, 0xa7, 0x87, 0x62, 0x96, 0x5e, 0x00, 0x54, 0x71, 0x00, 0x58, 0xad,
	0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0x5f, 0x01, 0x8a, 0xa0, 0x90, 0x24, 0x17, 0x47, 0x8e, 0x79,
	0x3c, 0xd8, 0x55, 0x12, 0xcc, 0x60, 0xfb, 0xd8, 0x73, 0xcc, 0x03, 0xc0, 0x8e, 0x94, 0xe7, 0xe2,
	0x2e, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0x8c, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x01, 0xcb, 0x72,
	0x41, 0x84, 0xfc, 0x12, 0x73, 0x53, 0xa5, 0x1c, 0xb9, 0x84, 0xb1, 0x58, 0x21, 0x24, 0xc0, 0xc5,
	0x9c, 0x9d, 0x5a, 0x09, 0x75, 0x3d, 0x88, 0x29, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a,
	0x2a, 0xc1, 0x04, 0x16, 0x83, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x93, 0xd8, 0xc0, 0x36, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x2e, 0x16, 0xe0, 0x35, 0x01, 0x00, 0x00,
}

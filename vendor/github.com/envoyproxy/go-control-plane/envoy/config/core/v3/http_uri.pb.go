// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/core/v3/http_uri.proto

package envoy_config_core_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_4026d4ca9f4f1539, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.config.core.v3.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/config/core/v3/http_uri.proto", fileDescriptor_4026d4ca9f4f1539)
}

var fileDescriptor_4026d4ca9f4f1539 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x6a, 0x32, 0x31,
	0x14, 0x85, 0x8d, 0xca, 0x3f, 0x7f, 0xd3, 0xae, 0x86, 0x42, 0x1d, 0x0b, 0x32, 0xd5, 0x8d, 0xab,
	0x04, 0xb4, 0xab, 0x6e, 0x0a, 0xa1, 0x0b, 0x97, 0x22, 0x74, 0x2d, 0x51, 0xe3, 0x34, 0xa0, 0xb9,
	0x21, 0x73, 0x13, 0xea, 0x1b, 0xf4, 0x19, 0xfa, 0x08, 0x7d, 0x8f, 0xd2, 0x57, 0x2a, 0xae, 0xca,
	0x4c, 0x66, 0x16, 0x85, 0xee, 0xc2, 0x3d, 0xdf, 0x09, 0x1f, 0x87, 0x4e, 0x94, 0x09, 0x70, 0xe2,
	0x5b, 0x30, 0x7b, 0x5d, 0xf0, 0x2d, 0x38, 0xc5, 0xc3, 0x9c, 0xbf, 0x20, 0xda, 0xb5, 0x77, 0x9a,
	0x59, 0x07, 0x08, 0xe9, 0x75, 0x0d, 0xb1, 0x08, 0xb1, 0x0a, 0x62, 0x61, 0x3e, 0x1c, 0x15, 0x00,
	0xc5, 0x41, 0xf1, 0x9a, 0xd9, 0xf8, 0x3d, 0xdf, 0x79, 0x27, 0x51, 0x83, 0x89, 0xad, 0xe1, 0x9d,
	0xdf, 0x59, 0xc9, 0xa5, 0x31, 0x80, 0xf5, 0xb9, 0xe4, 0x41, 0xb9, 0x52, 0x83, 0xd1, 0xa6, 0x68,
	0x90, 0x9b, 0x20, 0x0f, 0x7a, 0x27, 0x51, 0xf1, 0xf6, 0x11, 0x83, 0xf1, 0x17, 0xa1, 0xc9, 0x02,
	0xd1, 0x3e, 0x3b, 0x9d, 0x66, 0xb4, 0xe7, 0x9d, 0x1e, 0x90, 0x9c, 0x4c, 0x2f, 0x44, 0x72, 0x16,
	0x7d, 0xd7, 0xcd, 0xc9, 0xaa, 0xba, 0xa5, 0x13, 0x9a, 0x6c, 0x0f, 0xbe, 0x44, 0xe5, 0x06, 0xdd,
	0x5f, 0xf1, 0xa2, 0xb3, 0x6a, 0x93, 0xf4, 0x91, 0x26, 0xa8, 0x8f, 0x0a, 0x3c, 0x0e, 0x7a, 0x39,
	0x99, 0x5e, 0xce, 0x32, 0x16, 0xcd, 0x59, 0x6b, 0xce, 0x9e, 0x1a, 0x73, 0x41, 0xcf, 0x22, 0xf9,
	0x20, 0xfd, 0xff, 0x64, 0xd6, 0x59, 0xb5, 0xad, 0x87, 0xfc, 0xfd, 0xf3, 0x6d, 0x74, 0x4b, 0xb3,
	0xb8, 0x82, 0xb4, 0x9a, 0x85, 0x59, 0x5c, 0xa1, 0x51, 0x14, 0x19, 0x4d, 0xe3, 0x64, 0xb6, 0x44,
	0xa7, 0xe4, 0x71, 0x8d, 0x27, 0xab, 0xd2, 0xde, 0xb7, 0x20, 0xe2, 0x9e, 0x8e, 0x35, 0xb0, 0xba,
	0x6a, 0x1d, 0xbc, 0x9e, 0xd8, 0x5f, 0x5b, 0x8a, 0xab, 0xe6, 0xa7, 0x65, 0x65, 0xb4, 0x24, 0x9b,
	0x7f, 0xb5, 0xda, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xbb, 0x99, 0x33, 0x9f, 0x01, 0x00,
	0x00,
}

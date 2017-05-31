// Code generated by protoc-gen-go.
// source: sample.proto
// DO NOT EDIT!

/*
Package testprotos is a generated protocol buffer package.

It is generated from these files:
	sample.proto

It has these top-level messages:
	SimpleMsg
	NestedMsg
*/
package testprotos

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

// SimpleMsg is designed to test that all three types of message fields, plain, map,
// and slice are handled by the protolator tool
type SimpleMsg struct {
	PlainField string            `protobuf:"bytes,1,opt,name=plain_field,json=plainField" json:"plain_field,omitempty"`
	MapField   map[string]string `protobuf:"bytes,2,rep,name=map_field,json=mapField" json:"map_field,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SliceField []string          `protobuf:"bytes,3,rep,name=slice_field,json=sliceField" json:"slice_field,omitempty"`
}

func (m *SimpleMsg) Reset()                    { *m = SimpleMsg{} }
func (m *SimpleMsg) String() string            { return proto.CompactTextString(m) }
func (*SimpleMsg) ProtoMessage()               {}
func (*SimpleMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SimpleMsg) GetMapField() map[string]string {
	if m != nil {
		return m.MapField
	}
	return nil
}

// NestedMsg is designed to test the nested message component
type NestedMsg struct {
	PlainNestedField *SimpleMsg            `protobuf:"bytes,1,opt,name=plain_nested_field,json=plainNestedField" json:"plain_nested_field,omitempty"`
	MapNestedField   map[string]*SimpleMsg `protobuf:"bytes,2,rep,name=map_nested_field,json=mapNestedField" json:"map_nested_field,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	SliceNestedField []*SimpleMsg          `protobuf:"bytes,3,rep,name=slice_nested_field,json=sliceNestedField" json:"slice_nested_field,omitempty"`
}

func (m *NestedMsg) Reset()                    { *m = NestedMsg{} }
func (m *NestedMsg) String() string            { return proto.CompactTextString(m) }
func (*NestedMsg) ProtoMessage()               {}
func (*NestedMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NestedMsg) GetPlainNestedField() *SimpleMsg {
	if m != nil {
		return m.PlainNestedField
	}
	return nil
}

func (m *NestedMsg) GetMapNestedField() map[string]*SimpleMsg {
	if m != nil {
		return m.MapNestedField
	}
	return nil
}

func (m *NestedMsg) GetSliceNestedField() []*SimpleMsg {
	if m != nil {
		return m.SliceNestedField
	}
	return nil
}

func init() {
	proto.RegisterType((*SimpleMsg)(nil), "testprotos.SimpleMsg")
	proto.RegisterType((*NestedMsg)(nil), "testprotos.NestedMsg")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x69, 0xcb, 0xff, 0x8f, 0x7d, 0xa7, 0x32, 0xaa, 0xc2, 0xd8, 0xc5, 0x31, 0x2f, 0x13,
	0xa1, 0x81, 0x79, 0x11, 0xbd, 0x8c, 0x89, 0xde, 0xe6, 0x61, 0xbb, 0x88, 0x17, 0xc9, 0xba, 0x6c,
	0x0b, 0x26, 0x4d, 0x48, 0x32, 0x61, 0xdf, 0xcf, 0xa3, 0x1f, 0x4a, 0xf2, 0xa6, 0x6c, 0x2d, 0x6c,
	0xa7, 0x26, 0x4f, 0xfa, 0xfe, 0x9e, 0xe7, 0x79, 0xe1, 0xd4, 0x52, 0xa9, 0x05, 0xcb, 0xb5, 0x51,
	0x4e, 0x65, 0xe0, 0x98, 0x75, 0x78, 0xb4, 0xfd, 0xdf, 0x08, 0xd2, 0x19, 0xf7, 0x8f, 0x13, 0xbb,
	0xca, 0xae, 0xa1, 0xa5, 0x05, 0xe5, 0xe5, 0xe7, 0x92, 0x33, 0xb1, 0xe8, 0x44, 0xbd, 0x68, 0x90,
	0x4e, 0x01, 0xa5, 0x57, 0xaf, 0x64, 0x23, 0x48, 0x25, 0xd5, 0xd5, 0x73, 0xdc, 0x4b, 0x06, 0xad,
	0xe1, 0x4d, 0xbe, 0xc7, 0xe5, 0x3b, 0x54, 0x3e, 0xa1, 0x1a, 0x47, 0x5e, 0x4a, 0x67, 0xb6, 0xd3,
	0x13, 0x59, 0x5d, 0xbd, 0x85, 0x15, 0xbc, 0x60, 0x15, 0x23, 0xe9, 0x25, 0xde, 0x02, 0x25, 0xfc,
	0xa1, 0xfb, 0x04, 0x67, 0x8d, 0xd9, 0xac, 0x0d, 0xc9, 0x17, 0xdb, 0x56, 0x61, 0xfc, 0x31, 0xbb,
	0x84, 0x7f, 0xdf, 0x54, 0x6c, 0x58, 0x27, 0x46, 0x2d, 0x5c, 0x1e, 0xe3, 0x87, 0xa8, 0xff, 0x13,
	0x43, 0xfa, 0xc6, 0xac, 0x63, 0x0b, 0x5f, 0xe7, 0x19, 0xb2, 0x50, 0xa7, 0x44, 0xa9, 0xd6, 0xaa,
	0x35, 0xbc, 0x3a, 0x18, 0x7b, 0xda, 0xc6, 0x81, 0x80, 0x08, 0x81, 0x67, 0xd0, 0xf6, 0x95, 0x1b,
	0x88, 0xd0, 0xfc, 0xb6, 0x8e, 0xd8, 0xb9, 0xfa, 0xe6, 0xb5, 0xf9, 0xd0, 0xff, 0x5c, 0x36, 0x44,
	0x9f, 0x2c, 0x6c, 0xa1, 0x81, 0x4d, 0x10, 0x7b, 0x2c, 0x19, 0x0e, 0xd4, 0x20, 0xdd, 0x77, 0xb8,
	0x38, 0xe0, 0x75, 0x60, 0x5f, 0x77, 0xf5, 0x7d, 0x1d, 0x35, 0xd8, 0xaf, 0x71, 0x3c, 0xfe, 0x18,
	0xad, 0xb8, 0x5b, 0x6f, 0xe6, 0x79, 0xa1, 0x24, 0x59, 0x6f, 0x35, 0x33, 0x82, 0x2d, 0x56, 0xcc,
	0x90, 0x25, 0x9d, 0x1b, 0x5e, 0x90, 0x42, 0x49, 0xa9, 0x4a, 0xe2, 0x94, 0x12, 0x96, 0x20, 0x49,
	0x50, 0xa7, 0x0c, 0xd9, 0x83, 0xe7, 0xff, 0xf1, 0x7b, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0x1e, 0xfa, 0x9e, 0x7c, 0x02, 0x00, 0x00,
}
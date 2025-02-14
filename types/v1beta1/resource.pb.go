// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta1/resource.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// CPU stores resource units and cpu config attributes
type CPU struct {
	Units      ResourceValue `protobuf:"bytes,1,opt,name=units,proto3" json:"units"`
	Attributes []Attribute   `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" yaml:"cpu,omitempty"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_c018ae61a1e1b59b, []int{0}
}
func (m *CPU) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return m.Size()
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetUnits() ResourceValue {
	if m != nil {
		return m.Units
	}
	return ResourceValue{}
}

func (m *CPU) GetAttributes() []Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Memory stores resource quantity and memory attributes
type Memory struct {
	Quantity   ResourceValue `protobuf:"bytes,1,opt,name=quantity,proto3" json:"size" yaml:"size"`
	Attributes []Attribute   `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" yaml:"cpu,omitempty"`
}

func (m *Memory) Reset()         { *m = Memory{} }
func (m *Memory) String() string { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()    {}
func (*Memory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c018ae61a1e1b59b, []int{1}
}
func (m *Memory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Memory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Memory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Memory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Memory.Merge(m, src)
}
func (m *Memory) XXX_Size() int {
	return m.Size()
}
func (m *Memory) XXX_DiscardUnknown() {
	xxx_messageInfo_Memory.DiscardUnknown(m)
}

var xxx_messageInfo_Memory proto.InternalMessageInfo

func (m *Memory) GetQuantity() ResourceValue {
	if m != nil {
		return m.Quantity
	}
	return ResourceValue{}
}

func (m *Memory) GetAttributes() []Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Storage stores resource quantity and storage attributes
type Storage struct {
	Quantity   ResourceValue `protobuf:"bytes,1,opt,name=quantity,proto3" json:"size" yaml:"size"`
	Attributes []Attribute   `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" yaml:"cpu,omitempty"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_c018ae61a1e1b59b, []int{2}
}
func (m *Storage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return m.Size()
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

func (m *Storage) GetQuantity() ResourceValue {
	if m != nil {
		return m.Quantity
	}
	return ResourceValue{}
}

func (m *Storage) GetAttributes() []Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// ResourceUnits describes all available resources types for deployment/node etc
// if field is nil resource is not present in the given data-structure
type ResourceUnits struct {
	CPU       *CPU       `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory    *Memory    `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty" yaml:"memory,omitempty"`
	Storage   *Storage   `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty" yaml:"storage,omitempty"`
	Endpoints []Endpoint `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints" yaml:"endpoints"`
}

func (m *ResourceUnits) Reset()         { *m = ResourceUnits{} }
func (m *ResourceUnits) String() string { return proto.CompactTextString(m) }
func (*ResourceUnits) ProtoMessage()    {}
func (*ResourceUnits) Descriptor() ([]byte, []int) {
	return fileDescriptor_c018ae61a1e1b59b, []int{3}
}
func (m *ResourceUnits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceUnits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceUnits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceUnits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUnits.Merge(m, src)
}
func (m *ResourceUnits) XXX_Size() int {
	return m.Size()
}
func (m *ResourceUnits) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUnits.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUnits proto.InternalMessageInfo

func (m *ResourceUnits) GetCPU() *CPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *ResourceUnits) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUnits) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ResourceUnits) GetEndpoints() []Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "akash.base.v1beta1.CPU")
	proto.RegisterType((*Memory)(nil), "akash.base.v1beta1.Memory")
	proto.RegisterType((*Storage)(nil), "akash.base.v1beta1.Storage")
	proto.RegisterType((*ResourceUnits)(nil), "akash.base.v1beta1.ResourceUnits")
}

func init() { proto.RegisterFile("akash/base/v1beta1/resource.proto", fileDescriptor_c018ae61a1e1b59b) }

var fileDescriptor_c018ae61a1e1b59b = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0x9a, 0x75, 0xab, 0xb3, 0x14, 0x6a, 0x58, 0x68, 0xd8, 0x6a, 0xa6, 0x1d, 0x50,
	0x0a, 0x6a, 0x42, 0x5b, 0xbd, 0x14, 0x3c, 0x98, 0xd2, 0xa3, 0x50, 0x22, 0xeb, 0xa1, 0x17, 0x99,
	0x6c, 0x87, 0x34, 0x74, 0x93, 0x89, 0x99, 0x49, 0x25, 0x7e, 0x0a, 0x3f, 0x82, 0x1f, 0xc1, 0x6f,
	0xe0, 0xb5, 0xc7, 0x1e, 0x3d, 0xc5, 0xb2, 0x7b, 0x91, 0x3d, 0xee, 0x27, 0x90, 0xcc, 0x4c, 0x76,
	0xdd, 0x1a, 0xc5, 0x9b, 0xd0, 0x5b, 0x32, 0xef, 0xff, 0xde, 0xef, 0x3d, 0xfe, 0x6f, 0x06, 0x6e,
	0x93, 0x73, 0xc2, 0xcf, 0xdc, 0x80, 0x70, 0xea, 0x5e, 0xec, 0x06, 0x54, 0x90, 0x5d, 0x37, 0xa3,
	0x9c, 0xe5, 0xd9, 0x90, 0x3a, 0x69, 0xc6, 0x04, 0x33, 0x4d, 0x29, 0x71, 0x2a, 0x89, 0xa3, 0x25,
	0xfd, 0x5e, 0xc8, 0x42, 0x26, 0xc3, 0x6e, 0xf5, 0xa5, 0x94, 0x7d, 0xdc, 0x50, 0x8c, 0x08, 0x91,
	0x45, 0x41, 0x2e, 0x74, 0xb5, 0xfe, 0xe3, 0xbf, 0x00, 0x2f, 0xc8, 0x28, 0xaf, 0x75, 0x4d, 0x8d,
	0xd1, 0xe4, 0x34, 0x65, 0x51, 0x22, 0x94, 0x04, 0x7f, 0x05, 0xd0, 0x38, 0x3c, 0x1e, 0x98, 0x2f,
	0xe1, 0x9d, 0x3c, 0x89, 0x04, 0xb7, 0xc0, 0x16, 0xd8, 0xe9, 0xee, 0x6d, 0x3b, 0xbf, 0x37, 0xec,
	0xf8, 0x1a, 0xf1, 0xb6, 0x42, 0x78, 0xed, 0xcb, 0x12, 0xb5, 0x7c, 0x95, 0x65, 0x72, 0x08, 0xe7,
	0x4d, 0x72, 0x6b, 0x65, 0xcb, 0xd8, 0xe9, 0xee, 0x3d, 0x6c, 0xaa, 0xf1, 0xaa, 0x56, 0x79, 0xcf,
	0xab, 0xfc, 0x69, 0x89, 0x7a, 0x8b, 0xc4, 0xa7, 0x2c, 0x8e, 0x04, 0x8d, 0x53, 0x51, 0xcc, 0x4a,
	0xd4, 0x2b, 0x48, 0x3c, 0x3a, 0xc0, 0xc3, 0x34, 0x5f, 0x1c, 0x63, 0xff, 0x17, 0xcc, 0x41, 0xfb,
	0xc7, 0x67, 0x04, 0xf0, 0x77, 0x00, 0x3b, 0xaf, 0x69, 0xcc, 0xb2, 0xc2, 0x3c, 0x81, 0x77, 0xdf,
	0xe7, 0x24, 0x11, 0x91, 0x28, 0xfe, 0x7d, 0x8e, 0x4d, 0xdd, 0x47, 0x9b, 0x47, 0x1f, 0xe9, 0xac,
	0x44, 0x5d, 0xc5, 0xad, 0xfe, 0xb0, 0x3f, 0xaf, 0xf7, 0x3f, 0x27, 0xbc, 0x06, 0x70, 0xf5, 0x8d,
	0x60, 0x19, 0x09, 0xe9, 0x6d, 0x1d, 0xf1, 0x8b, 0x01, 0xd7, 0xea, 0x9e, 0x07, 0x72, 0xa3, 0xde,
	0x41, 0x63, 0x98, 0xe6, 0x7a, 0xc6, 0x8d, 0xa6, 0x2e, 0x0e, 0x8f, 0x07, 0x92, 0x0f, 0xc6, 0x25,
	0xaa, 0x76, 0x78, 0x5a, 0xa2, 0xb5, 0x25, 0xd0, 0x1f, 0xf9, 0x55, 0x65, 0x33, 0x84, 0x9d, 0x58,
	0xae, 0x8d, 0xb5, 0x22, 0x19, 0xfd, 0x26, 0x86, 0x5a, 0x2c, 0x6f, 0xbf, 0xc2, 0x4c, 0x4b, 0xb4,
	0xae, 0x32, 0x96, 0x10, 0x1b, 0x0a, 0x71, 0x33, 0x82, 0x7d, 0x5d, 0xde, 0x1c, 0xc1, 0x55, 0xae,
	0xdc, 0xb3, 0x0c, 0x49, 0xda, 0x6c, 0x22, 0x69, 0x83, 0xbd, 0x17, 0x1a, 0x75, 0x5f, 0xe7, 0x2c,
	0xb1, 0x2c, 0x6d, 0xdc, 0xcd, 0x10, 0xf6, 0x6b, 0x84, 0x49, 0xe0, 0xbd, 0xfa, 0x8a, 0x73, 0xab,
	0x2d, 0x3d, 0x7c, 0xd0, 0xc4, 0x3b, 0xd2, 0x22, 0xef, 0x91, 0xb6, 0x70, 0x91, 0x36, 0x2b, 0xd1,
	0xba, 0x02, 0xcd, 0x8f, 0xb0, 0xbf, 0x08, 0x2b, 0xcb, 0xbc, 0xa3, 0xcb, 0xb1, 0x0d, 0xae, 0xc6,
	0x36, 0xb8, 0x1e, 0xdb, 0xe0, 0xd3, 0xc4, 0x6e, 0x5d, 0x4d, 0xec, 0xd6, 0xb7, 0x89, 0xdd, 0x3a,
	0x79, 0x12, 0x46, 0xe2, 0x2c, 0x0f, 0x9c, 0x21, 0x8b, 0x5d, 0x49, 0x7e, 0x96, 0x50, 0xf1, 0x81,
	0x65, 0xe7, 0x6e, 0xc2, 0x4e, 0xa9, 0x2b, 0x8a, 0x94, 0xf2, 0xfa, 0x3d, 0x0a, 0x3a, 0xf2, 0x1d,
	0xda, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x00, 0x9a, 0xa1, 0xe0, 0x45, 0x05, 0x00, 0x00,
}

func (this *CPU) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CPU)
	if !ok {
		that2, ok := that.(CPU)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Units.Equal(&that1.Units) {
		return false
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(&that1.Attributes[i]) {
			return false
		}
	}
	return true
}
func (this *Memory) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Memory)
	if !ok {
		that2, ok := that.(Memory)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Quantity.Equal(&that1.Quantity) {
		return false
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(&that1.Attributes[i]) {
			return false
		}
	}
	return true
}
func (this *Storage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Storage)
	if !ok {
		that2, ok := that.(Storage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Quantity.Equal(&that1.Quantity) {
		return false
	}
	if len(this.Attributes) != len(that1.Attributes) {
		return false
	}
	for i := range this.Attributes {
		if !this.Attributes[i].Equal(&that1.Attributes[i]) {
			return false
		}
	}
	return true
}
func (this *ResourceUnits) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceUnits)
	if !ok {
		that2, ok := that.(ResourceUnits)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.CPU.Equal(that1.CPU) {
		return false
	}
	if !this.Memory.Equal(that1.Memory) {
		return false
	}
	if !this.Storage.Equal(that1.Storage) {
		return false
	}
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(&that1.Endpoints[i]) {
			return false
		}
	}
	return true
}
func (m *CPU) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CPU) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CPU) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Units.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Memory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Memory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Memory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Quantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Storage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Storage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Storage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		for iNdEx := len(m.Attributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Quantity.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintResource(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ResourceUnits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceUnits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceUnits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResource(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Storage != nil {
		{
			size, err := m.Storage.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Memory != nil {
		{
			size, err := m.Memory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CPU != nil {
		{
			size, err := m.CPU.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResource(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResource(dAtA []byte, offset int, v uint64) int {
	offset -= sovResource(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CPU) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Units.Size()
	n += 1 + l + sovResource(uint64(l))
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovResource(uint64(l))
		}
	}
	return n
}

func (m *Memory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Quantity.Size()
	n += 1 + l + sovResource(uint64(l))
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovResource(uint64(l))
		}
	}
	return n
}

func (m *Storage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Quantity.Size()
	n += 1 + l + sovResource(uint64(l))
	if len(m.Attributes) > 0 {
		for _, e := range m.Attributes {
			l = e.Size()
			n += 1 + l + sovResource(uint64(l))
		}
	}
	return n
}

func (m *ResourceUnits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CPU != nil {
		l = m.CPU.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	if m.Memory != nil {
		l = m.Memory.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	if m.Storage != nil {
		l = m.Storage.Size()
		n += 1 + l + sovResource(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovResource(uint64(l))
		}
	}
	return n
}

func sovResource(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResource(x uint64) (n int) {
	return sovResource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CPU) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CPU: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CPU: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Units.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Memory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Memory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Memory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Storage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Storage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Storage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes, Attribute{})
			if err := m.Attributes[len(m.Attributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ResourceUnits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ResourceUnits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceUnits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CPU == nil {
				m.CPU = &CPU{}
			}
			if err := m.CPU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Memory == nil {
				m.Memory = &Memory{}
			}
			if err := m.Memory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Storage == nil {
				m.Storage = &Storage{}
			}
			if err := m.Storage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthResource
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResource
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipResource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResource
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowResource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthResource
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResource
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResource
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResource        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResource          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResource = fmt.Errorf("proto: unexpected end of group")
)

// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: uber/cadence/api/v1/visibility.proto

package apiv1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type IndexedValueType int32

const (
	IndexedValueType_INDEXED_VALUE_TYPE_INVALID  IndexedValueType = 0
	IndexedValueType_INDEXED_VALUE_TYPE_STRING   IndexedValueType = 1
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD  IndexedValueType = 2
	IndexedValueType_INDEXED_VALUE_TYPE_INT      IndexedValueType = 3
	IndexedValueType_INDEXED_VALUE_TYPE_DOUBLE   IndexedValueType = 4
	IndexedValueType_INDEXED_VALUE_TYPE_BOOL     IndexedValueType = 5
	IndexedValueType_INDEXED_VALUE_TYPE_DATETIME IndexedValueType = 6
)

var IndexedValueType_name = map[int32]string{
	0: "INDEXED_VALUE_TYPE_INVALID",
	1: "INDEXED_VALUE_TYPE_STRING",
	2: "INDEXED_VALUE_TYPE_KEYWORD",
	3: "INDEXED_VALUE_TYPE_INT",
	4: "INDEXED_VALUE_TYPE_DOUBLE",
	5: "INDEXED_VALUE_TYPE_BOOL",
	6: "INDEXED_VALUE_TYPE_DATETIME",
}

var IndexedValueType_value = map[string]int32{
	"INDEXED_VALUE_TYPE_INVALID":  0,
	"INDEXED_VALUE_TYPE_STRING":   1,
	"INDEXED_VALUE_TYPE_KEYWORD":  2,
	"INDEXED_VALUE_TYPE_INT":      3,
	"INDEXED_VALUE_TYPE_DOUBLE":   4,
	"INDEXED_VALUE_TYPE_BOOL":     5,
	"INDEXED_VALUE_TYPE_DATETIME": 6,
}

func (x IndexedValueType) String() string {
	return proto.EnumName(IndexedValueType_name, int32(x))
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1b2132ef24217c8, []int{0}
}

type WorkflowExecutionFilter struct {
	WorkflowId           string   `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId                string   `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowExecutionFilter) Reset()         { *m = WorkflowExecutionFilter{} }
func (m *WorkflowExecutionFilter) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionFilter) ProtoMessage()    {}
func (*WorkflowExecutionFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b2132ef24217c8, []int{0}
}
func (m *WorkflowExecutionFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowExecutionFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowExecutionFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowExecutionFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionFilter.Merge(m, src)
}
func (m *WorkflowExecutionFilter) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowExecutionFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionFilter.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionFilter proto.InternalMessageInfo

func (m *WorkflowExecutionFilter) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowExecutionFilter) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type WorkflowTypeFilter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkflowTypeFilter) Reset()         { *m = WorkflowTypeFilter{} }
func (m *WorkflowTypeFilter) String() string { return proto.CompactTextString(m) }
func (*WorkflowTypeFilter) ProtoMessage()    {}
func (*WorkflowTypeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b2132ef24217c8, []int{1}
}
func (m *WorkflowTypeFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowTypeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowTypeFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowTypeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTypeFilter.Merge(m, src)
}
func (m *WorkflowTypeFilter) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowTypeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTypeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTypeFilter proto.InternalMessageInfo

func (m *WorkflowTypeFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StartTimeFilter struct {
	EarliestTime         *types.Timestamp `protobuf:"bytes,1,opt,name=earliest_time,json=earliestTime,proto3" json:"earliest_time,omitempty"`
	LatestTime           *types.Timestamp `protobuf:"bytes,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StartTimeFilter) Reset()         { *m = StartTimeFilter{} }
func (m *StartTimeFilter) String() string { return proto.CompactTextString(m) }
func (*StartTimeFilter) ProtoMessage()    {}
func (*StartTimeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b2132ef24217c8, []int{2}
}
func (m *StartTimeFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartTimeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartTimeFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartTimeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartTimeFilter.Merge(m, src)
}
func (m *StartTimeFilter) XXX_Size() int {
	return m.Size()
}
func (m *StartTimeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StartTimeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StartTimeFilter proto.InternalMessageInfo

func (m *StartTimeFilter) GetEarliestTime() *types.Timestamp {
	if m != nil {
		return m.EarliestTime
	}
	return nil
}

func (m *StartTimeFilter) GetLatestTime() *types.Timestamp {
	if m != nil {
		return m.LatestTime
	}
	return nil
}

type StatusFilter struct {
	Status               WorkflowExecutionCloseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=uber.cadence.api.v1.WorkflowExecutionCloseStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StatusFilter) Reset()         { *m = StatusFilter{} }
func (m *StatusFilter) String() string { return proto.CompactTextString(m) }
func (*StatusFilter) ProtoMessage()    {}
func (*StatusFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b2132ef24217c8, []int{3}
}
func (m *StatusFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusFilter.Merge(m, src)
}
func (m *StatusFilter) XXX_Size() int {
	return m.Size()
}
func (m *StatusFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StatusFilter proto.InternalMessageInfo

func (m *StatusFilter) GetStatus() WorkflowExecutionCloseStatus {
	if m != nil {
		return m.Status
	}
	return WorkflowExecutionCloseStatus_WORKFLOW_EXECUTION_CLOSE_STATUS_INVALID
}

func init() {
	proto.RegisterEnum("uber.cadence.api.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
	proto.RegisterType((*WorkflowExecutionFilter)(nil), "uber.cadence.api.v1.WorkflowExecutionFilter")
	proto.RegisterType((*WorkflowTypeFilter)(nil), "uber.cadence.api.v1.WorkflowTypeFilter")
	proto.RegisterType((*StartTimeFilter)(nil), "uber.cadence.api.v1.StartTimeFilter")
	proto.RegisterType((*StatusFilter)(nil), "uber.cadence.api.v1.StatusFilter")
}

func init() {
	proto.RegisterFile("uber/cadence/api/v1/visibility.proto", fileDescriptor_c1b2132ef24217c8)
}

var fileDescriptor_c1b2132ef24217c8 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x71, 0x68, 0x23, 0xb8, 0x29, 0x60, 0x0d, 0x82, 0x80, 0x2b, 0x12, 0x64, 0xb1, 0xa8,
	0x58, 0x8c, 0x95, 0xb2, 0xec, 0x02, 0x25, 0xd8, 0xa0, 0x11, 0x21, 0x09, 0x8e, 0x9b, 0x12, 0x36,
	0xd6, 0xd8, 0x9e, 0x86, 0x11, 0xb6, 0xc7, 0xb2, 0xc7, 0x69, 0xfb, 0x14, 0xbc, 0x16, 0x4b, 0x1e,
	0x01, 0x65, 0xc9, 0x53, 0x20, 0xff, 0x21, 0x21, 0x1c, 0xb1, 0xb3, 0xcf, 0x3d, 0xe7, 0xd3, 0xdc,
	0x1f, 0x78, 0x91, 0x7b, 0x2c, 0x35, 0x7c, 0x1a, 0xb0, 0xd8, 0x67, 0x06, 0x4d, 0xb8, 0xb1, 0x1d,
	0x19, 0x5b, 0x9e, 0x71, 0x8f, 0x87, 0x5c, 0xde, 0xe0, 0x24, 0x15, 0x52, 0xa0, 0x87, 0x85, 0x0b,
	0xd7, 0x2e, 0x4c, 0x13, 0x8e, 0xb7, 0x23, 0x6d, 0xb8, 0x11, 0x62, 0x13, 0x32, 0xa3, 0xb4, 0x78,
	0xf9, 0xa5, 0x21, 0x79, 0xc4, 0x32, 0x49, 0xa3, 0xa4, 0x4a, 0x69, 0x7a, 0x1b, 0xfb, 0x4a, 0xa4,
	0x5f, 0x2f, 0x43, 0x71, 0x55, 0x79, 0xf4, 0x8f, 0xd0, 0xbf, 0xa8, 0x15, 0xeb, 0x9a, 0xf9, 0xb9,
	0xe4, 0x22, 0x7e, 0xcb, 0x43, 0xc9, 0x52, 0x34, 0x84, 0x5e, 0x63, 0x76, 0x79, 0xf0, 0x44, 0x79,
	0xae, 0x9c, 0xdc, 0xb5, 0xa1, 0x91, 0x48, 0x80, 0x1e, 0x41, 0x37, 0xcd, 0xe3, 0xa2, 0xd6, 0x29,
	0x6b, 0x87, 0x69, 0x1e, 0x93, 0x40, 0x3f, 0x01, 0xd4, 0x20, 0x9d, 0x9b, 0x84, 0xd5, 0x34, 0x04,
	0x07, 0x31, 0x8d, 0x58, 0x8d, 0x29, 0xbf, 0xf5, 0x6f, 0x0a, 0x3c, 0x58, 0x4a, 0x9a, 0x4a, 0x87,
	0x47, 0x8d, 0xef, 0x35, 0xdc, 0x63, 0x34, 0x0d, 0x39, 0xcb, 0xa4, 0x5b, 0x34, 0x54, 0x06, 0x7a,
	0xa7, 0x1a, 0xae, 0xba, 0xc5, 0x4d, 0xb7, 0xd8, 0x69, 0xba, 0xb5, 0x8f, 0x9a, 0x40, 0x21, 0xa1,
	0x33, 0xe8, 0x85, 0x54, 0xfe, 0x89, 0x77, 0xfe, 0x1b, 0x87, 0xca, 0x5e, 0x08, 0xfa, 0x1a, 0x8e,
	0x96, 0x92, 0xca, 0x3c, 0xab, 0x5f, 0x43, 0xa0, 0x9b, 0x95, 0xff, 0xe5, 0x33, 0xee, 0x9f, 0x8e,
	0x70, 0xcb, 0x26, 0xf0, 0x3f, 0x13, 0x7c, 0x13, 0x8a, 0x8c, 0x55, 0x20, 0xbb, 0x06, 0xbc, 0xfc,
	0xa5, 0x80, 0x4a, 0xe2, 0x80, 0x5d, 0xb3, 0x60, 0x45, 0xc3, 0x9c, 0x15, 0xb3, 0x41, 0x03, 0xd0,
	0xc8, 0xcc, 0xb4, 0x3e, 0x59, 0xa6, 0xbb, 0x1a, 0x4f, 0xcf, 0x2d, 0xd7, 0x59, 0x2f, 0x2c, 0x97,
	0xcc, 0x56, 0xe3, 0x29, 0x31, 0xd5, 0x5b, 0xe8, 0x19, 0x3c, 0x6d, 0xa9, 0x2f, 0x1d, 0x9b, 0xcc,
	0xde, 0xa9, 0xca, 0x9e, 0xf8, 0x7b, 0x6b, 0x7d, 0x31, 0xb7, 0x4d, 0xb5, 0x83, 0x34, 0x78, 0xdc,
	0x8a, 0x77, 0xd4, 0xdb, 0x7b, 0xd0, 0xe6, 0xfc, 0x7c, 0x32, 0xb5, 0xd4, 0x03, 0x74, 0x0c, 0xfd,
	0x96, 0xf2, 0x64, 0x3e, 0x9f, 0xaa, 0x87, 0x68, 0x08, 0xc7, 0x6d, 0xd9, 0xb1, 0x63, 0x39, 0xe4,
	0x83, 0xa5, 0x76, 0x27, 0xde, 0xf7, 0xdd, 0x40, 0xf9, 0xb1, 0x1b, 0x28, 0x3f, 0x77, 0x03, 0x05,
	0xfa, 0xbe, 0x88, 0xda, 0x06, 0x37, 0xb9, 0x33, 0x4e, 0xf8, 0xa2, 0xd8, 0xc8, 0x42, 0xf9, 0x6c,
	0x6c, 0xb8, 0xfc, 0x92, 0x7b, 0xd8, 0x17, 0x91, 0xf1, 0xd7, 0xe1, 0xe2, 0x0d, 0x8b, 0xab, 0x23,
	0xaf, 0x6f, 0xf8, 0x8c, 0x26, 0x7c, 0x3b, 0xf2, 0xba, 0xa5, 0xf6, 0xea, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe8, 0x16, 0x6d, 0xcb, 0x43, 0x03, 0x00, 0x00,
}

func (m *WorkflowExecutionFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowExecutionFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowExecutionFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintVisibility(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintVisibility(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkflowTypeFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowTypeFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowTypeFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintVisibility(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StartTimeFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartTimeFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StartTimeFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.LatestTime != nil {
		{
			size, err := m.LatestTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVisibility(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.EarliestTime != nil {
		{
			size, err := m.EarliestTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVisibility(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StatusFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintVisibility(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVisibility(dAtA []byte, offset int, v uint64) int {
	offset -= sovVisibility(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowExecutionFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovVisibility(uint64(l))
	}
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovVisibility(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WorkflowTypeFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovVisibility(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StartTimeFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EarliestTime != nil {
		l = m.EarliestTime.Size()
		n += 1 + l + sovVisibility(uint64(l))
	}
	if m.LatestTime != nil {
		l = m.LatestTime.Size()
		n += 1 + l + sovVisibility(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StatusFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovVisibility(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovVisibility(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVisibility(x uint64) (n int) {
	return sovVisibility(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkflowExecutionFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVisibility
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
			return fmt.Errorf("proto: WorkflowExecutionFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowExecutionFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVisibility
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVisibility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVisibility
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVisibility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVisibility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVisibility
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WorkflowTypeFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVisibility
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
			return fmt.Errorf("proto: WorkflowTypeFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowTypeFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthVisibility
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVisibility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVisibility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVisibility
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StartTimeFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVisibility
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
			return fmt.Errorf("proto: StartTimeFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartTimeFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarliestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
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
				return ErrInvalidLengthVisibility
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVisibility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EarliestTime == nil {
				m.EarliestTime = &types.Timestamp{}
			}
			if err := m.EarliestTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
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
				return ErrInvalidLengthVisibility
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVisibility
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LatestTime == nil {
				m.LatestTime = &types.Timestamp{}
			}
			if err := m.LatestTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVisibility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVisibility
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StatusFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVisibility
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
			return fmt.Errorf("proto: StatusFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVisibility
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= WorkflowExecutionCloseStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVisibility(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVisibility
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVisibility(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVisibility
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
					return 0, ErrIntOverflowVisibility
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
					return 0, ErrIntOverflowVisibility
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
				return 0, ErrInvalidLengthVisibility
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVisibility
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVisibility
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVisibility        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVisibility          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVisibility = fmt.Errorf("proto: unexpected end of group")
)

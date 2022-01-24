// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/dsl.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type FindOneReq struct {
	TableName            string     `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Dsl                  *anypb.Any `protobuf:"bytes,2,opt,name=dsl,proto3" json:"dsl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindOneReq) Reset()         { *m = FindOneReq{} }
func (m *FindOneReq) String() string { return proto.CompactTextString(m) }
func (*FindOneReq) ProtoMessage()    {}
func (*FindOneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{0}
}

func (m *FindOneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneReq.Unmarshal(m, b)
}
func (m *FindOneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneReq.Marshal(b, m, deterministic)
}
func (m *FindOneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneReq.Merge(m, src)
}
func (m *FindOneReq) XXX_Size() int {
	return xxx_messageInfo_FindOneReq.Size(m)
}
func (m *FindOneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneReq proto.InternalMessageInfo

func (m *FindOneReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *FindOneReq) GetDsl() *anypb.Any {
	if m != nil {
		return m.Dsl
	}
	return nil
}

type FindOneResp struct {
	Data                 *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindOneResp) Reset()         { *m = FindOneResp{} }
func (m *FindOneResp) String() string { return proto.CompactTextString(m) }
func (*FindOneResp) ProtoMessage()    {}
func (*FindOneResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{1}
}

func (m *FindOneResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOneResp.Unmarshal(m, b)
}
func (m *FindOneResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOneResp.Marshal(b, m, deterministic)
}
func (m *FindOneResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOneResp.Merge(m, src)
}
func (m *FindOneResp) XXX_Size() int {
	return xxx_messageInfo_FindOneResp.Size(m)
}
func (m *FindOneResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOneResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindOneResp proto.InternalMessageInfo

func (m *FindOneResp) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

type FindReq struct {
	TableName            string     `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Page                 int64      `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64      `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Sort                 []string   `protobuf:"bytes,4,rep,name=sort,proto3" json:"sort,omitempty"`
	Dsl                  *anypb.Any `protobuf:"bytes,5,opt,name=dsl,proto3" json:"dsl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindReq) Reset()         { *m = FindReq{} }
func (m *FindReq) String() string { return proto.CompactTextString(m) }
func (*FindReq) ProtoMessage()    {}
func (*FindReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{2}
}

func (m *FindReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindReq.Unmarshal(m, b)
}
func (m *FindReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindReq.Marshal(b, m, deterministic)
}
func (m *FindReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindReq.Merge(m, src)
}
func (m *FindReq) XXX_Size() int {
	return xxx_messageInfo_FindReq.Size(m)
}
func (m *FindReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindReq proto.InternalMessageInfo

func (m *FindReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *FindReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FindReq) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *FindReq) GetSort() []string {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *FindReq) GetDsl() *anypb.Any {
	if m != nil {
		return m.Dsl
	}
	return nil
}

type FindResp struct {
	Data                 *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Count                int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindResp) Reset()         { *m = FindResp{} }
func (m *FindResp) String() string { return proto.CompactTextString(m) }
func (*FindResp) ProtoMessage()    {}
func (*FindResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{3}
}

func (m *FindResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindResp.Unmarshal(m, b)
}
func (m *FindResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindResp.Marshal(b, m, deterministic)
}
func (m *FindResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindResp.Merge(m, src)
}
func (m *FindResp) XXX_Size() int {
	return xxx_messageInfo_FindResp.Size(m)
}
func (m *FindResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindResp proto.InternalMessageInfo

func (m *FindResp) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FindResp) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CountReq struct {
	TableName            string     `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Dsl                  *anypb.Any `protobuf:"bytes,2,opt,name=dsl,proto3" json:"dsl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CountReq) Reset()         { *m = CountReq{} }
func (m *CountReq) String() string { return proto.CompactTextString(m) }
func (*CountReq) ProtoMessage()    {}
func (*CountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{4}
}

func (m *CountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountReq.Unmarshal(m, b)
}
func (m *CountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountReq.Marshal(b, m, deterministic)
}
func (m *CountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountReq.Merge(m, src)
}
func (m *CountReq) XXX_Size() int {
	return xxx_messageInfo_CountReq.Size(m)
}
func (m *CountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CountReq.DiscardUnknown(m)
}

var xxx_messageInfo_CountReq proto.InternalMessageInfo

func (m *CountReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *CountReq) GetDsl() *anypb.Any {
	if m != nil {
		return m.Dsl
	}
	return nil
}

type CountResp struct {
	Data                 int64    `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountResp) Reset()         { *m = CountResp{} }
func (m *CountResp) String() string { return proto.CompactTextString(m) }
func (*CountResp) ProtoMessage()    {}
func (*CountResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{5}
}

func (m *CountResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountResp.Unmarshal(m, b)
}
func (m *CountResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountResp.Marshal(b, m, deterministic)
}
func (m *CountResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountResp.Merge(m, src)
}
func (m *CountResp) XXX_Size() int {
	return xxx_messageInfo_CountResp.Size(m)
}
func (m *CountResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CountResp.DiscardUnknown(m)
}

var xxx_messageInfo_CountResp proto.InternalMessageInfo

func (m *CountResp) GetData() int64 {
	if m != nil {
		return m.Data
	}
	return 0
}

type InsertReq struct {
	TableName            string       `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Entities             []*anypb.Any `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *InsertReq) Reset()         { *m = InsertReq{} }
func (m *InsertReq) String() string { return proto.CompactTextString(m) }
func (*InsertReq) ProtoMessage()    {}
func (*InsertReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{6}
}

func (m *InsertReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertReq.Unmarshal(m, b)
}
func (m *InsertReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertReq.Marshal(b, m, deterministic)
}
func (m *InsertReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertReq.Merge(m, src)
}
func (m *InsertReq) XXX_Size() int {
	return xxx_messageInfo_InsertReq.Size(m)
}
func (m *InsertReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertReq proto.InternalMessageInfo

func (m *InsertReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *InsertReq) GetEntities() []*anypb.Any {
	if m != nil {
		return m.Entities
	}
	return nil
}

type InsertResp struct {
	Count                int64    `protobuf:"zigzag64,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertResp) Reset()         { *m = InsertResp{} }
func (m *InsertResp) String() string { return proto.CompactTextString(m) }
func (*InsertResp) ProtoMessage()    {}
func (*InsertResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{7}
}

func (m *InsertResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertResp.Unmarshal(m, b)
}
func (m *InsertResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertResp.Marshal(b, m, deterministic)
}
func (m *InsertResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertResp.Merge(m, src)
}
func (m *InsertResp) XXX_Size() int {
	return xxx_messageInfo_InsertResp.Size(m)
}
func (m *InsertResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertResp.DiscardUnknown(m)
}

var xxx_messageInfo_InsertResp proto.InternalMessageInfo

func (m *InsertResp) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type UpdateReq struct {
	TableName            string     `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Dsl                  *anypb.Any `protobuf:"bytes,2,opt,name=dsl,proto3" json:"dsl,omitempty"`
	Entity               *anypb.Any `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{8}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *UpdateReq) GetDsl() *anypb.Any {
	if m != nil {
		return m.Dsl
	}
	return nil
}

func (m *UpdateReq) GetEntity() *anypb.Any {
	if m != nil {
		return m.Entity
	}
	return nil
}

type UpdateResp struct {
	Count                int64    `protobuf:"zigzag64,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResp) Reset()         { *m = UpdateResp{} }
func (m *UpdateResp) String() string { return proto.CompactTextString(m) }
func (*UpdateResp) ProtoMessage()    {}
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{9}
}

func (m *UpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResp.Unmarshal(m, b)
}
func (m *UpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResp.Marshal(b, m, deterministic)
}
func (m *UpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResp.Merge(m, src)
}
func (m *UpdateResp) XXX_Size() int {
	return xxx_messageInfo_UpdateResp.Size(m)
}
func (m *UpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResp proto.InternalMessageInfo

func (m *UpdateResp) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DeleteReq struct {
	TableName            string     `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Dsl                  *anypb.Any `protobuf:"bytes,2,opt,name=dsl,proto3" json:"dsl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DeleteReq) Reset()         { *m = DeleteReq{} }
func (m *DeleteReq) String() string { return proto.CompactTextString(m) }
func (*DeleteReq) ProtoMessage()    {}
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{10}
}

func (m *DeleteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteReq.Unmarshal(m, b)
}
func (m *DeleteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteReq.Marshal(b, m, deterministic)
}
func (m *DeleteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteReq.Merge(m, src)
}
func (m *DeleteReq) XXX_Size() int {
	return xxx_messageInfo_DeleteReq.Size(m)
}
func (m *DeleteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteReq proto.InternalMessageInfo

func (m *DeleteReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *DeleteReq) GetDsl() *anypb.Any {
	if m != nil {
		return m.Dsl
	}
	return nil
}

type DeleteResp struct {
	Count                int64    `protobuf:"zigzag64,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResp) Reset()         { *m = DeleteResp{} }
func (m *DeleteResp) String() string { return proto.CompactTextString(m) }
func (*DeleteResp) ProtoMessage()    {}
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_279353885e3c64ca, []int{11}
}

func (m *DeleteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResp.Unmarshal(m, b)
}
func (m *DeleteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResp.Marshal(b, m, deterministic)
}
func (m *DeleteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResp.Merge(m, src)
}
func (m *DeleteResp) XXX_Size() int {
	return xxx_messageInfo_DeleteResp.Size(m)
}
func (m *DeleteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResp proto.InternalMessageInfo

func (m *DeleteResp) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*FindOneReq)(nil), "proto.FindOneReq")
	proto.RegisterType((*FindOneResp)(nil), "proto.FindOneResp")
	proto.RegisterType((*FindReq)(nil), "proto.FindReq")
	proto.RegisterType((*FindResp)(nil), "proto.FindResp")
	proto.RegisterType((*CountReq)(nil), "proto.CountReq")
	proto.RegisterType((*CountResp)(nil), "proto.CountResp")
	proto.RegisterType((*InsertReq)(nil), "proto.InsertReq")
	proto.RegisterType((*InsertResp)(nil), "proto.InsertResp")
	proto.RegisterType((*UpdateReq)(nil), "proto.UpdateReq")
	proto.RegisterType((*UpdateResp)(nil), "proto.UpdateResp")
	proto.RegisterType((*DeleteReq)(nil), "proto.DeleteReq")
	proto.RegisterType((*DeleteResp)(nil), "proto.DeleteResp")
}

func init() { proto.RegisterFile("api/proto/dsl.proto", fileDescriptor_279353885e3c64ca) }

var fileDescriptor_279353885e3c64ca = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x25, 0xcd, 0xc7, 0x6b, 0xee, 0x03, 0xf5, 0x5d, 0xdf, 0x22, 0x06, 0xc1, 0x90, 0x85, 0x06,
	0xf1, 0xa5, 0x52, 0x17, 0xae, 0xc5, 0x22, 0x28, 0xe2, 0x47, 0x8a, 0x2b, 0x57, 0xd3, 0x66, 0x2c,
	0x81, 0x98, 0x8c, 0x99, 0xa9, 0x50, 0x37, 0xfe, 0x00, 0xff, 0x84, 0x3f, 0x55, 0x66, 0x26, 0x33,
	0x69, 0x0b, 0xa1, 0x3e, 0xe8, 0x2a, 0x77, 0xce, 0x9c, 0x9c, 0x7b, 0xee, 0xc7, 0xc0, 0x7d, 0xc2,
	0xaa, 0x19, 0xeb, 0x5a, 0xd1, 0xce, 0x4a, 0x5e, 0xe7, 0x2a, 0x42, 0x5f, 0x7d, 0xe2, 0x07, 0x9b,
	0xb6, 0xdd, 0xd4, 0x54, 0x5f, 0xaf, 0xb6, 0xdf, 0x66, 0xa4, 0xd9, 0x69, 0x46, 0x5a, 0x00, 0xbc,
	0xa9, 0x9a, 0xf2, 0x63, 0x43, 0x0b, 0xfa, 0x03, 0x1f, 0x42, 0x28, 0xc8, 0xaa, 0xa6, 0x1f, 0xc8,
	0x77, 0x1a, 0x39, 0x89, 0x93, 0x85, 0xc5, 0x00, 0xe0, 0x63, 0x70, 0x4b, 0x5e, 0x47, 0x93, 0xc4,
	0xc9, 0x2e, 0xe7, 0xd7, 0xb9, 0x16, 0xcd, 0x8d, 0x68, 0xfe, 0xaa, 0xd9, 0x15, 0x92, 0x90, 0xbe,
	0x84, 0x4b, 0xab, 0xc9, 0x19, 0x66, 0xe0, 0x95, 0x44, 0x10, 0xa5, 0x37, 0xf6, 0x9f, 0x62, 0xa4,
	0x7f, 0x1c, 0xb8, 0x90, 0x7f, 0x9e, 0xb6, 0x82, 0xe0, 0x31, 0xb2, 0xa1, 0xca, 0x8b, 0x5b, 0xa8,
	0x58, 0x62, 0xbc, 0xfa, 0x45, 0x23, 0x57, 0x63, 0x32, 0x56, 0x58, 0xdb, 0x89, 0xc8, 0x4b, 0xdc,
	0x2c, 0x2c, 0x54, 0x6c, 0xca, 0xf0, 0x4f, 0x95, 0xf1, 0x0e, 0xa6, 0xda, 0xcc, 0x6d, 0x6a, 0xc0,
	0x6b, 0xf0, 0xd7, 0xed, 0xb6, 0x11, 0xbd, 0x35, 0x7d, 0x48, 0x3f, 0xc1, 0xf4, 0xb5, 0x0c, 0xce,
	0xd7, 0xe4, 0x47, 0x10, 0xf6, 0x8a, 0x9c, 0xc9, 0x32, 0xad, 0x3d, 0xb7, 0x6f, 0xe6, 0x57, 0x08,
	0xdf, 0x36, 0x9c, 0x76, 0xff, 0x91, 0xf3, 0x39, 0x4c, 0x69, 0x23, 0x2a, 0x51, 0x51, 0x1e, 0x4d,
	0x12, 0x77, 0x34, 0xb1, 0x65, 0xa5, 0x29, 0x80, 0x11, 0xe7, 0x6c, 0xa8, 0x59, 0x2a, 0xa3, 0xa9,
	0xf9, 0x37, 0x84, 0x5f, 0x58, 0x49, 0xc4, 0xf9, 0x36, 0x0b, 0x9f, 0x41, 0xa0, 0x2c, 0xec, 0xd4,
	0x90, 0xc7, 0xa8, 0x3d, 0x47, 0x9a, 0x34, 0x06, 0x46, 0x4d, 0x7e, 0x86, 0x70, 0x41, 0x6b, 0x7a,
	0x46, 0x93, 0x32, 0xad, 0x91, 0x1c, 0x4b, 0x3b, 0xff, 0x3b, 0x01, 0x58, 0x2c, 0xdf, 0x2f, 0x69,
	0xf7, 0xb3, 0x5a, 0xcb, 0x01, 0x5c, 0xf4, 0x2f, 0x06, 0xaf, 0xb4, 0x62, 0x3e, 0xbc, 0xca, 0x18,
	0x8f, 0x21, 0xce, 0xf0, 0x09, 0x78, 0xf2, 0x88, 0x77, 0xf6, 0xee, 0x24, 0xf7, 0xee, 0xc1, 0x99,
	0x33, 0x7c, 0x0a, 0xbe, 0xda, 0x13, 0x34, 0x37, 0x66, 0x0f, 0xe3, 0x7b, 0x87, 0x00, 0x67, 0x78,
	0x03, 0x81, 0x9e, 0x2a, 0x9a, 0x3b, 0xbb, 0x41, 0xf1, 0xd5, 0x11, 0xa2, 0xe9, 0xba, 0xbf, 0x96,
	0x6e, 0xe7, 0x6d, 0xe9, 0x7b, 0x03, 0xb8, 0x81, 0x40, 0xf7, 0xc5, 0xd2, 0x6d, 0xe7, 0x2d, 0x7d,
	0x68, 0xdc, 0x2a, 0x50, 0xc8, 0x8b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x7e, 0x58, 0xd7,
	0xd9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DSLServiceClient is the client API for DSLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DSLServiceClient interface {
	FindOne(ctx context.Context, in *FindOneReq, opts ...grpc.CallOption) (*FindOneResp, error)
	Find(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindResp, error)
	Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountResp, error)
	Insert(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertResp, error)
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error)
}

type dSLServiceClient struct {
	cc *grpc.ClientConn
}

func NewDSLServiceClient(cc *grpc.ClientConn) DSLServiceClient {
	return &dSLServiceClient{cc}
}

func (c *dSLServiceClient) FindOne(ctx context.Context, in *FindOneReq, opts ...grpc.CallOption) (*FindOneResp, error) {
	out := new(FindOneResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLServiceClient) Find(ctx context.Context, in *FindReq, opts ...grpc.CallOption) (*FindResp, error) {
	out := new(FindResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLServiceClient) Count(ctx context.Context, in *CountReq, opts ...grpc.CallOption) (*CountResp, error) {
	out := new(CountResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLServiceClient) Insert(ctx context.Context, in *InsertReq, opts ...grpc.CallOption) (*InsertResp, error) {
	out := new(InsertResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLServiceClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateResp, error) {
	out := new(UpdateResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLServiceClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteResp, error) {
	out := new(DeleteResp)
	err := c.cc.Invoke(ctx, "/proto.DSLService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DSLServiceServer is the server API for DSLService service.
type DSLServiceServer interface {
	FindOne(context.Context, *FindOneReq) (*FindOneResp, error)
	Find(context.Context, *FindReq) (*FindResp, error)
	Count(context.Context, *CountReq) (*CountResp, error)
	Insert(context.Context, *InsertReq) (*InsertResp, error)
	Update(context.Context, *UpdateReq) (*UpdateResp, error)
	Delete(context.Context, *DeleteReq) (*DeleteResp, error)
}

// UnimplementedDSLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDSLServiceServer struct {
}

func (*UnimplementedDSLServiceServer) FindOne(ctx context.Context, req *FindOneReq) (*FindOneResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (*UnimplementedDSLServiceServer) Find(ctx context.Context, req *FindReq) (*FindResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (*UnimplementedDSLServiceServer) Count(ctx context.Context, req *CountReq) (*CountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (*UnimplementedDSLServiceServer) Insert(ctx context.Context, req *InsertReq) (*InsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (*UnimplementedDSLServiceServer) Update(ctx context.Context, req *UpdateReq) (*UpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedDSLServiceServer) Delete(ctx context.Context, req *DeleteReq) (*DeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterDSLServiceServer(s *grpc.Server, srv DSLServiceServer) {
	s.RegisterService(&_DSLService_serviceDesc, srv)
}

func _DSLService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).FindOne(ctx, req.(*FindOneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).Find(ctx, req.(*FindReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).Count(ctx, req.(*CountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).Insert(ctx, req.(*InsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLServiceServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DSLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DSLService",
	HandlerType: (*DSLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOne",
			Handler:    _DSLService_FindOne_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _DSLService_Find_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _DSLService_Count_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _DSLService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DSLService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DSLService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/dsl.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/ddl.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ExecuteReq struct {
	TableName            string   `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	Option               string   `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
	Fields               []*Field `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteReq) Reset()         { *m = ExecuteReq{} }
func (m *ExecuteReq) String() string { return proto.CompactTextString(m) }
func (*ExecuteReq) ProtoMessage()    {}
func (*ExecuteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bd5749e671052d, []int{0}
}

func (m *ExecuteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteReq.Unmarshal(m, b)
}
func (m *ExecuteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteReq.Marshal(b, m, deterministic)
}
func (m *ExecuteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteReq.Merge(m, src)
}
func (m *ExecuteReq) XXX_Size() int {
	return xxx_messageInfo_ExecuteReq.Size(m)
}
func (m *ExecuteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteReq.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteReq proto.InternalMessageInfo

func (m *ExecuteReq) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *ExecuteReq) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *ExecuteReq) GetFields() []*Field {
	if m != nil {
		return m.Fields
	}
	return nil
}

type ExecuteResp struct {
	TableName            string   `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteResp) Reset()         { *m = ExecuteResp{} }
func (m *ExecuteResp) String() string { return proto.CompactTextString(m) }
func (*ExecuteResp) ProtoMessage()    {}
func (*ExecuteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bd5749e671052d, []int{1}
}

func (m *ExecuteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResp.Unmarshal(m, b)
}
func (m *ExecuteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResp.Marshal(b, m, deterministic)
}
func (m *ExecuteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResp.Merge(m, src)
}
func (m *ExecuteResp) XXX_Size() int {
	return xxx_messageInfo_ExecuteResp.Size(m)
}
func (m *ExecuteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResp.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResp proto.InternalMessageInfo

func (m *ExecuteResp) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

type Field struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	NotNull              bool     `protobuf:"varint,4,opt,name=notNull,proto3" json:"notNull,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9bd5749e671052d, []int{2}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Field) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Field) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Field) GetNotNull() bool {
	if m != nil {
		return m.NotNull
	}
	return false
}

func init() {
	proto.RegisterType((*ExecuteReq)(nil), "proto.ExecuteReq")
	proto.RegisterType((*ExecuteResp)(nil), "proto.ExecuteResp")
	proto.RegisterType((*Field)(nil), "proto.Field")
}

func init() { proto.RegisterFile("api/proto/ddl.proto", fileDescriptor_c9bd5749e671052d) }

var fileDescriptor_c9bd5749e671052d = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0xdd, 0x76, 0xdd, 0x59, 0x2f, 0x8e, 0x22, 0x41, 0x3c, 0x94, 0xe2, 0xa1, 0x20,
	0xec, 0xca, 0x7a, 0xf7, 0xb4, 0x7a, 0x92, 0x1e, 0xe2, 0x27, 0xe8, 0x9f, 0x11, 0x03, 0x69, 0x13,
	0xdb, 0xa9, 0xe8, 0xb7, 0x97, 0xa6, 0xa9, 0x05, 0x0f, 0x9e, 0x32, 0xef, 0xf7, 0x12, 0xde, 0xcb,
	0xc0, 0x45, 0x61, 0xd5, 0xde, 0x76, 0x86, 0xcd, 0xbe, 0xae, 0xf5, 0xce, 0x4d, 0x18, 0xb9, 0x23,
	0x7d, 0x07, 0x78, 0xfa, 0xa2, 0x6a, 0x60, 0x92, 0xf4, 0x81, 0x37, 0xb0, 0xe1, 0xa2, 0xd4, 0x94,
	0x17, 0x0d, 0x89, 0x20, 0x09, 0xb2, 0x8d, 0x5c, 0x00, 0x5e, 0x41, 0x6c, 0x2c, 0x2b, 0xd3, 0x8a,
	0x13, 0x67, 0x79, 0x85, 0xb7, 0x10, 0xbf, 0x29, 0xd2, 0x75, 0x2f, 0xc2, 0x24, 0xcc, 0xb6, 0x87,
	0xb3, 0x29, 0x62, 0xf7, 0x3c, 0x42, 0xe9, 0xbd, 0xf4, 0x0e, 0xb6, 0xbf, 0x49, 0xbd, 0xfd, 0x3f,
	0x2a, 0x25, 0x88, 0xdc, 0x6b, 0xbc, 0x84, 0x88, 0x15, 0xeb, 0xf9, 0xca, 0x24, 0x10, 0x61, 0xc5,
	0xdf, 0x96, 0x7c, 0x0f, 0x37, 0xa3, 0x80, 0x75, 0x65, 0x9a, 0x86, 0x5a, 0x16, 0xa1, 0xc3, 0xb3,
	0x1c, 0x9d, 0xd6, 0x70, 0x3e, 0x68, 0x2d, 0x56, 0x49, 0x90, 0x9d, 0xca, 0x59, 0x1e, 0x1e, 0x01,
	0x8e, 0xc7, 0x97, 0x57, 0xea, 0x3e, 0x55, 0x45, 0x78, 0x0f, 0x6b, 0xdf, 0x10, 0xcf, 0xfd, 0x17,
	0x96, 0xdd, 0x5c, 0xe3, 0x5f, 0xd4, 0xdb, 0x32, 0x76, 0xe8, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x7b, 0xfe, 0x5b, 0x3e, 0x62, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DDLServiceClient is the client API for DDLService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DDLServiceClient interface {
	Execute(ctx context.Context, in *ExecuteReq, opts ...grpc.CallOption) (*ExecuteResp, error)
}

type dDLServiceClient struct {
	cc *grpc.ClientConn
}

func NewDDLServiceClient(cc *grpc.ClientConn) DDLServiceClient {
	return &dDLServiceClient{cc}
}

func (c *dDLServiceClient) Execute(ctx context.Context, in *ExecuteReq, opts ...grpc.CallOption) (*ExecuteResp, error) {
	out := new(ExecuteResp)
	err := c.cc.Invoke(ctx, "/proto.DDLService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DDLServiceServer is the server API for DDLService service.
type DDLServiceServer interface {
	Execute(context.Context, *ExecuteReq) (*ExecuteResp, error)
}

// UnimplementedDDLServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDDLServiceServer struct {
}

func (*UnimplementedDDLServiceServer) Execute(ctx context.Context, req *ExecuteReq) (*ExecuteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterDDLServiceServer(s *grpc.Server, srv DDLServiceServer) {
	s.RegisterService(&_DDLService_serviceDesc, srv)
}

func _DDLService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDLServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DDLService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDLServiceServer).Execute(ctx, req.(*ExecuteReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DDLService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DDLService",
	HandlerType: (*DDLServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _DDLService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/ddl.proto",
}
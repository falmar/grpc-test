// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Todo struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Completed            bool     `protobuf:"varint,3,opt,name=Completed,proto3" json:"Completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_db92bbe3fa49d516, []int{0}
}
func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (dst *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(dst, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Todo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Todo) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_db92bbe3fa49d516, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ListRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Completed            bool     `protobuf:"varint,4,opt,name=Completed,proto3" json:"Completed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_db92bbe3fa49d516, []int{2}
}
func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (dst *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(dst, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ListRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListRequest) GetCompleted() bool {
	if m != nil {
		return m.Completed
	}
	return false
}

type GetRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_todo_db92bbe3fa49d516, []int{3}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func init() {
	proto.RegisterType((*Todo)(nil), "pb.Todo")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*ListRequest)(nil), "pb.ListRequest")
	proto.RegisterType((*GetRequest)(nil), "pb.GetRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TODOClient is the client API for TODO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TODOClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TODO_ListClient, error)
	Create(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	MarkCompleted(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error)
	Delete(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Empty, error)
}

type tODOClient struct {
	cc *grpc.ClientConn
}

func NewTODOClient(cc *grpc.ClientConn) TODOClient {
	return &tODOClient{cc}
}

func (c *tODOClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (TODO_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TODO_serviceDesc.Streams[0], "/pb.TODO/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &tODOListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TODO_ListClient interface {
	Recv() (*Todo, error)
	grpc.ClientStream
}

type tODOListClient struct {
	grpc.ClientStream
}

func (x *tODOListClient) Recv() (*Todo, error) {
	m := new(Todo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tODOClient) Create(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/pb.TODO/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tODOClient) MarkCompleted(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Todo, error) {
	out := new(Todo)
	err := c.cc.Invoke(ctx, "/pb.TODO/MarkCompleted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tODOClient) Delete(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.TODO/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TODOServer is the server API for TODO service.
type TODOServer interface {
	List(*ListRequest, TODO_ListServer) error
	Create(context.Context, *Todo) (*Todo, error)
	MarkCompleted(context.Context, *Todo) (*Todo, error)
	Delete(context.Context, *Todo) (*Empty, error)
}

func RegisterTODOServer(s *grpc.Server, srv TODOServer) {
	s.RegisterService(&_TODO_serviceDesc, srv)
}

func _TODO_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TODOServer).List(m, &tODOListServer{stream})
}

type TODO_ListServer interface {
	Send(*Todo) error
	grpc.ServerStream
}

type tODOListServer struct {
	grpc.ServerStream
}

func (x *tODOListServer) Send(m *Todo) error {
	return x.ServerStream.SendMsg(m)
}

func _TODO_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TODOServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TODO/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TODOServer).Create(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TODO_MarkCompleted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TODOServer).MarkCompleted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TODO/MarkCompleted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TODOServer).MarkCompleted(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TODO_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TODOServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TODO/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TODOServer).Delete(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

var _TODO_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TODO",
	HandlerType: (*TODOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TODO_Create_Handler,
		},
		{
			MethodName: "MarkCompleted",
			Handler:    _TODO_MarkCompleted_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TODO_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _TODO_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "todo.proto",
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_todo_db92bbe3fa49d516) }

var fileDescriptor_todo_db92bbe3fa49d516 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x51, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0xc9, 0x36, 0x36, 0x23, 0x2a, 0x0c, 0x22, 0x41, 0x8a, 0x96, 0x40, 0xa1, 0x4f,
	0x41, 0xf4, 0x08, 0x8d, 0x68, 0xa1, 0x1a, 0x5c, 0x7a, 0x81, 0x84, 0x4c, 0x21, 0x68, 0xd8, 0x35,
	0x99, 0x3e, 0xf4, 0x20, 0xde, 0x57, 0x76, 0x53, 0x62, 0x2c, 0xf4, 0x6d, 0xfe, 0x7f, 0x7f, 0xe6,
	0x9b, 0x9d, 0x01, 0x60, 0x5d, 0xea, 0xc4, 0x34, 0x9a, 0x35, 0x7a, 0xa6, 0x88, 0x5f, 0x41, 0x6e,
	0x74, 0xa9, 0xf1, 0x12, 0xbc, 0x55, 0x1a, 0x89, 0x99, 0x58, 0x84, 0xca, 0x5b, 0xa5, 0x88, 0x20,
	0xdf, 0xf3, 0x9a, 0x22, 0xcf, 0x39, 0xae, 0xc6, 0x29, 0x84, 0x4b, 0x5d, 0x9b, 0x2f, 0x62, 0x2a,
	0x23, 0x7f, 0x26, 0x16, 0x13, 0xf5, 0x67, 0xc4, 0x67, 0x30, 0x7e, 0xae, 0x0d, 0xef, 0x63, 0x0d,
	0xe7, 0xeb, 0xaa, 0x65, 0x45, 0xdf, 0x3b, 0x6a, 0x19, 0xaf, 0x61, 0xfc, 0xb1, 0xa3, 0x66, 0x7f,
	0x68, 0xde, 0x09, 0xeb, 0xae, 0xab, 0xba, 0x62, 0x07, 0xf0, 0x55, 0x27, 0xf0, 0x06, 0x82, 0x6c,
	0xbb, 0x6d, 0x89, 0x5d, 0x7b, 0x5f, 0x1d, 0xd4, 0x7f, 0xb2, 0x3c, 0x26, 0x4f, 0x01, 0x5e, 0xa8,
	0xe7, 0x1d, 0xfd, 0xe4, 0xf1, 0x47, 0x80, 0xdc, 0x64, 0x69, 0x86, 0x73, 0x90, 0x76, 0x2e, 0xbc,
	0x4a, 0x4c, 0x91, 0x0c, 0x26, 0xbc, 0x9d, 0x58, 0xc3, 0x6e, 0x21, 0x1e, 0x3d, 0x08, 0xbc, 0x83,
	0x60, 0xd9, 0x50, 0xce, 0x84, 0xbd, 0x3f, 0x4c, 0xe0, 0x1c, 0x2e, 0xde, 0xf2, 0xe6, 0xb3, 0xc7,
	0x9f, 0x88, 0xdd, 0x43, 0x90, 0x92, 0x0d, 0x0c, 0xde, 0x43, 0x5b, 0x75, 0x4b, 0x1a, 0x15, 0x81,
	0x3b, 0xc2, 0xd3, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x69, 0x37, 0x0f, 0x92, 0x01, 0x00,
	0x00,
}

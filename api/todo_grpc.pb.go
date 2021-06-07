// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	GetTasksStream(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (Todo_GetTasksStreamClient, error)
	GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksReply, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) GetTasksStream(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (Todo_GetTasksStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Todo_ServiceDesc.Streams[0], "/todo.Todo/GetTasksStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoGetTasksStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Todo_GetTasksStreamClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type todoGetTasksStreamClient struct {
	grpc.ClientStream
}

func (x *todoGetTasksStreamClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *todoClient) GetTasks(ctx context.Context, in *GetTasksRequest, opts ...grpc.CallOption) (*GetTasksReply, error) {
	out := new(GetTasksReply)
	err := c.cc.Invoke(ctx, "/todo.Todo/GetTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/todo.Todo/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	GetTasksStream(*GetTasksRequest, Todo_GetTasksStreamServer) error
	GetTasks(context.Context, *GetTasksRequest) (*GetTasksReply, error)
	CreateTask(context.Context, *CreateTaskRequest) (*Task, error)
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) GetTasksStream(*GetTasksRequest, Todo_GetTasksStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTasksStream not implemented")
}
func (UnimplementedTodoServer) GetTasks(context.Context, *GetTasksRequest) (*GetTasksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasks not implemented")
}
func (UnimplementedTodoServer) CreateTask(context.Context, *CreateTaskRequest) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_GetTasksStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTasksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServer).GetTasksStream(m, &todoGetTasksStreamServer{stream})
}

type Todo_GetTasksStreamServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type todoGetTasksStreamServer struct {
	grpc.ServerStream
}

func (x *todoGetTasksStreamServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _Todo_GetTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).GetTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/GetTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).GetTasks(ctx, req.(*GetTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Todo/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTasks",
			Handler:    _Todo_GetTasks_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Todo_CreateTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTasksStream",
			Handler:       _Todo_GetTasksStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/todo.proto",
}
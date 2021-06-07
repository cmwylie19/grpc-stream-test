package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/cmwylie19/grpc-tests/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Task struct {
	Name    string
	Created string
}

var tl []*api.Task

type server struct {
	api.UnimplementedTodoServer
}

func (s *server) CreateTask(ctx context.Context, in *api.CreateTaskRequest) (*api.Task, error) {
	t := Task{
		Name:    in.GetName(),
		Created: time.Now().String(),
	}
	tl = append(tl, &api.Task{
		Name:    in.GetName(),
		Created: time.Now().String(),
	})
	return &api.Task{
		Name:    t.Name,
		Created: t.Created,
	}, nil
}

func (s *server) GetTasks(ctx context.Context, in *api.GetTasksRequest) (*api.GetTasksReply, error) {
	fmt.Println("taskArr ", tl)
	return &api.GetTasksReply{
		Task: tl,
	}, nil
}

func (s *server) GetTasksStream(req *api.GetTasksRequest, stream api.Todo_GetTasksStreamServer) error {
	fmt.Printf("GetTasksStream function was invoked: %v", req)
	for i, s := range tl {
		fmt.Println(i, s)
		res := &api.Task{
			Name:    s.Name,
			Created: s.Created,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func taskUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// preprocessing logic
	// gets info about the current RPC call by examining the args passed in
	log.Println("========= [Server interceptor] ", info.FullMethod)

	// Invoke the handler to complete the normal execution of a unary RPC
	m, err := handler(ctx, req)

	//Post processing logic
	log.Printf("Post Proc Message: %s", m)
	return m, err
}

type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("========= [Server Stream Interceptor Wrapper - RcvMsg] "+"Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("========= [Server Stream Interceptor Wrapper - SendMsg] "+"Send a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func taskServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("========= [Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(taskUnaryServerInterceptor), grpc.StreamInterceptor(taskServerStreamInterceptor))
	api.RegisterTodoServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

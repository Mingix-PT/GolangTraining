package main

import (
    "context"
    "log"
    "net"

    pb "grpc-example/todo/todo"

    "google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct {
    pb.UnimplementedTodoServiceServer
}

func (s *server) GetTodo(ctx context.Context, req *pb.GetTodoRequest) (*pb.TodoResponse, error) {
    log.Printf("Received request for Todo ID: %v", req.Id)
    return &pb.TodoResponse{
        Id:    req.Id,
        Title: "this is a mock todo API",
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTodoServiceServer(s, &server{})
    log.Printf("Server started at %s", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

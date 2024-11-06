package main

import (
	"context"
	"log"
	"net"

	pb "GolangTraining/grpc/todo/todo"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type TodoServer struct {
	pb.UnimplementedTodoServer
}

func (s *TodoServer) GetTodo(ctx context.Context, in *pb.RequestTodo) (*pb.ReplyTodo, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.ReplyTodo{Id: in.GetId(), Msg: "This is a mock Todo."}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTodoServer(s, &TodoServer{})
	log.Printf("Server started on port %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

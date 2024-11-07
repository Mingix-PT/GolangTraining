package main

import (
	"log"
	"net"

	"context"

	"github.com/Mingix-PT/demo-grpc/gettodo"
	"google.golang.org/grpc"
)

type getToDoServer struct {
	gettodo.UnimplementedRandomToDoServer
}

func (s *getToDoServer) GetToDo(ctx context.Context, in *gettodo.GetToDoRequest) (*gettodo.GetToDoResponse, error) {
	return &gettodo.GetToDoResponse{Id: "1", Title: "This is a mock API"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}
	s := grpc.NewServer()
	service := &getToDoServer{}
	gettodo.RegisterRandomToDoServer(s, service)
	err = s.Serve(lis)
	log.Println("server is running on port 8080")
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

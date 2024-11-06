package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "GolangTraining/grpc/todo/todo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultId = 1
)

var (
	address = flag.String("address", "localhost:50051", "the address to connect to")
	id      = flag.Int("id", defaultId, "Request todo with id")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewTodoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetTodo(ctx, &pb.RequestTodo{Id: int32(*id)})
	if err != nil {
		log.Fatalf("Failed to get todo: %v", err)
	}

	log.Printf("Received: %v", r)
}

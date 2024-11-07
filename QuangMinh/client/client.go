package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Mingix-PT/demo-grpc/gettodo"
)

var (
	addr = flag.String("address", "localhost:8080", "address of the server")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRandomToDoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetToDo(ctx, &pb.GetToDoRequest{Id: "1"})
	if err != nil {
		log.Fatalf("could not get todo: %v", err)
	}

	log.Printf("Todo: %s", r.Title)
}

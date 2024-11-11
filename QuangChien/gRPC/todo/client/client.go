package main

import (
    "context"
    "log"
    "time"

    pb "grpc-example/todo/todo"

    "google.golang.org/grpc"
)

const (
    address = "localhost:50051"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    client := pb.NewTodoServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    reqId := int32(1) // Tùy chọn ID nào đó
    res, err := client.GetTodo(ctx, &pb.GetTodoRequest{Id: reqId})
    if err != nil {
        log.Fatalf("could not get todo: %v", err)
    }
    log.Printf("Todo: ID=%v, Title=%s", res.Id, res.Title)
}

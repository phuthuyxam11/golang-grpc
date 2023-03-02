package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_client.study/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1997", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error connect grpc server: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	// Code removed for brevity

	client := pb.NewTodoServiceClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	r, err := client.CreateTodo(context.Background(), &pb.CreateTodoRequest{
		Title:    "hello",
		Cat:      "melll",
		Auth:     nil,
		CreateAt: nil,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Response of grpc server: %s", r)
}

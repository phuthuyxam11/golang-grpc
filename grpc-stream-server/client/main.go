package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_stream_server.study/pb"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:1998", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	// Call the Stream method on the server using the client
	stream, err := client.CreateTodoStream(context.Background(), &pb.CreateTodoRequest{
		Title:    "",
		Cat:      "",
		Auth:     nil,
		CreateAt: nil,
	})
	if err != nil {
		log.Fatalf("failed to stream: %v", err)
	}

	// Receive streaming response from the server
	for {
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}
		log.Printf("received: %v", response.Title)
	}
}

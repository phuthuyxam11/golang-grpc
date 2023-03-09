package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"unary_pb.study/pb"
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//r, err := client.DeleteTodoWithDeadline(ctx, &pb.DeleteTodoRequest{
	//	Id:         -1,
	//	DeleteType: 0,
	//})

	r, err := client.DeleteTodoWithDeadline(ctx, &pb.DeleteTodoRequest{
		Id:         1,
		DeleteType: 0,
	})

	if err != nil {
		// ouch!
		// let's print the gRPC error message
		// which is "Length of `Name` cannot be more than 10 characters"
		errStatus, _ := status.FromError(err)

		if errStatus.Code() == codes.DeadlineExceeded {
			log.Println("server deadline")
		}

		fmt.Println(errStatus.Message())
		// let's print the error code which is `INVALID_ARGUMENT`
		fmt.Println(errStatus.Code())
		// Want its int version for some reason?
		// you shouldn't actually do this, but if you need for debugging,
		// you can do `int(status_code)` which will give you `3`
		//
		// Want to take specific action based on specific error?
		if codes.InvalidArgument == errStatus.Code() {
			// do your stuff here
			log.Fatal()
		}
	}

	log.Printf("Response of grpc server: %s", r)
}

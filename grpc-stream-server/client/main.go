package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc_stream_server.study/pb"
	"io/ioutil"
	"log"
	"time"
)

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	pemServerCA, err := ioutil.ReadFile("../ssl/ca-cert.pem")
	if err != nil {
		return nil, err
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("fail to add server CA's certificate")
	}
	config := &tls.Config{
		RootCAs: certPool,
	}
	return credentials.NewTLS(config), nil
}

func main() {
	credential, err := LoadTLSCredentials()
	if err != nil {
		log.Fatalln("Error create client credential: ", err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "localhost:1998", grpc.WithTransportCredentials(credential))
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

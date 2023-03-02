package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc.stream.client/pb"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// client stream
	filePath, _ := filepath.Abs("../tmp/csv_example.csv")
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("can't read file: ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.Dial("localhost:1999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error connect grpc server: ", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	client := pb.NewUploadServiceClient(conn)

	stream, err := client.UploadFileStream(ctx)
	if err != nil {
		fmt.Println("can't create stream: ", err)
		return
	}

	err = stream.Send(&pb.UploadFileRequest{Data: &pb.UploadFileRequest_Info{Info: &pb.ImageInfo{
		Filepath: "/tmp",
		Filetype: "csv",
		Filename: "example",
	}}})
	if err != nil {
		fmt.Println("can't send data to server")
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
		return
	}
	//chunk file
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}
		err = stream.Send(&pb.UploadFileRequest{Data: &pb.UploadFileRequest_Chunk{
			Chunk: buffer[:n],
		}})
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}
	fmt.Println("pre send")
	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	fmt.Println("respone: ", res)
}

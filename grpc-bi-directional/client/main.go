package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc.bi.directional/pb"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	// client stream
	filePath, _ := filepath.Abs("../tmp/csv_example.csv")
	file, err := os.Open(filePath)
	//waitc := make(chan struct{})
	defer file.Close()
	if err != nil {
		fmt.Println("can't read file: ", err)
		return
	}
	conn, err := grpc.Dial("localhost:1998", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Error connect grpc server: ", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	client := pb.NewSplitFileServiceClient(conn)

	stream, err := client.SplitFile(context.Background())
	if err != nil {
		fmt.Println("can't create stream: ", err)
		return
	}

	err = stream.Send(&pb.SplitFileRequest{Data: &pb.SplitFileRequest_FileInfo{FileInfo: &pb.FileInfo{
		FilePath: "/tmp",
		FileName: "csv_example",
	}}})
	if err != nil {
		fmt.Println("can't send data to server")
		log.Fatal("cannot send file info to server: ", err, stream.RecvMsg(nil))
		return
	}

	//chunk file
	rd := bufio.NewReader(file)
	var fileData [][]byte
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		fileData = append(fileData, []byte(line))
	}
	// send header &pb.SplitFileRequest_Chunk{fileData[0]}
	fmt.Println("len data", len(fileData))
	err = stream.Send(&pb.SplitFileRequest{Data: &pb.SplitFileRequest_Chunk{Chunk: fileData[0]}})
	if err != nil {
		log.Fatalln("Send header error: ", err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for k, v := range fileData {
			if k == 0 {
				continue
			}
			fmt.Println("count: ", k)
			err = stream.Send(&pb.SplitFileRequest{Data: &pb.SplitFileRequest_Chunk{Chunk: v}})
			if err != nil {
				fmt.Println("errororor", err)
				log.Fatalln("Send body error: ", err)
			}
		}
		//err := stream.CloseSend()
		//if err != nil {
		//	return
		//}
		wg.Done()
	}()

	fmt.Println("get ......")
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Fatalf("ending send server: %v", err)
				return
			}
			if err != nil {
				log.Fatalf("failed to recv: %v", err)
				return
			}

			log.Printf("filePath: %s, status: %s\n", resp.FilePath, resp.Status)
			if resp.Finish {
				break
			}
		}
		//close(waitc)
		wg.Done()
	}()
	//<-waitc
	wg.Wait()
	return
}

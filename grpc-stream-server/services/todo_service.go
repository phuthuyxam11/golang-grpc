package services

import (
	"encoding/csv"
	"fmt"
	"grpc_stream_server.study/pb"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Server) CreateTodoStream(request *pb.CreateTodoRequest, stream pb.TodoService_CreateTodoStreamServer) error {
	// open file
	path, err := filepath.Abs("../grpc-stream-server/assets/csv_example.csv")
	if err != nil {
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	count := 0
	for {
		time.Sleep(1 * time.Second)
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		response := pb.CreateTodoResponse{
			Title:    fmt.Sprintf("%d %s", count, strings.Join(rec, " ")),
			Cat:      "",
			Auth:     nil,
			CreateAt: nil,
		}
		if err := stream.Send(&response); err != nil {
			return err
		}
		count++
	}

	return nil
}

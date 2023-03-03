package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc.bi.directional/pb"
	"io"
	"log"
	"os"
	"path/filepath"
)

func (s *Server) SplitFile(stream pb.SplitFileService_SplitFileServer) error {
	fileStream, err := stream.Recv()
	if err != nil {
		log.Fatalln("Can't read file info csv: ", err)
		return err
	}
	fileInfo := fileStream.GetFileInfo()
	count := 0

	fmt.Println("file info: ", fileInfo)

	header, err := stream.Recv()
	if err != nil {
		log.Fatalln("Can't read header csv: ", err)
		return err
	}
	fmt.Println("file header: ", header)
	for {
		count++
		err := contextError(stream.Context())
		if err == io.EOF {
			// end of one request
			break
		}
		if err != nil {
			log.Fatalln("Error of context: ", err)
			return err
		}
		// file action
		absFilePath := fmt.Sprintf("./%s/%s", fileInfo.FilePath, fileInfo.FileName)
		filePath, _ := filepath.Abs(absFilePath)
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			log.Fatalln("Can't create folder split: ", err)
			return err
		}
		file, err := os.Create(fmt.Sprintf("%s/%s_%d.csv", filePath, fileInfo.FileName, count))
		if err != nil {
			log.Fatalln("Can't create file split: ", err)
			return err
		}
		_, err = file.Write(header.GetChunk())
		if err != nil {
			log.Fatalln("Can't write header: ", err)
			return err
		}

		// write content file
		data, err := stream.Recv()
		fmt.Println("file body: ", data)
		if err != nil {
			log.Fatalln("Can't read body file: ", err)
			return err
		}
		_, err = file.Write(data.GetChunk())
		err = file.Close()
		if err != nil {
			log.Fatalln("File can't close: ", err)
			return err
		}
		response := &pb.SplitFileResponse{
			Status:   "ok",
			FilePath: filePath,
			Finish:   false,
		}
		err = stream.Send(response)
		if err != nil {
			log.Fatalln("Can't send file to client: ", err)
			return err
		}
	}

	err = stream.Send(&pb.SplitFileResponse{
		Status:   "ok",
		FilePath: "",
		Finish:   true,
	})
	if err != nil {
		log.Fatalln("Can't finish message to client: ", err)
		return err
	}
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}

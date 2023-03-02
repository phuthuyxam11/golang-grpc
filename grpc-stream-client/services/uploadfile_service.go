package services

import (
	"bytes"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc.stream.client/pb"
	"io"
	"log"
)

func (s *Server) UploadFileStream(stream pb.UploadService_UploadFileStreamServer) error {
	req, err := stream.Recv()
	if err != nil {
		log.Fatalln("Can't receive request: ", err)
		return err
	}
	fileInfo := req.GetInfo()
	storeUpload := NewDiskStoreFile(fileInfo.Filepath)

	fileData := bytes.Buffer{}
	fileSize := 0
	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}
		req, err := stream.Recv()

		if err == io.EOF {
			log.Print("finish request")
			// receive success one request
			break
		}
		if err != nil {
			log.Fatalln("Has some bull-shit err: ", err)
			return err
		}
		chunkData := req.GetChunk()
		fileSize += len(chunkData)
		// limit file size if required
		_, err = fileData.Write(chunkData)
		if err != nil {
			log.Fatalln("Can not read byte data: ", err)
			return err
		}
	}
	err = storeUpload.Save(fileInfo.Filepath, fileInfo.Filename, fileInfo.Filetype, fileData)
	if err != nil {
		return err
	}

	// send response to client
	res := &pb.UploadFileResponse{
		Status:   "ok",
		FilePath: fileInfo.Filepath,
	}
	err = stream.SendAndClose(res)
	if err != nil {
		log.Fatalln("Can't send message to client: ", err)
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

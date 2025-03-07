package main

import (
	"log"
	"net"
	"os"

	file_store "github.com/Milindcode/file_store/grpc/proto"
	"google.golang.org/grpc"
)

type FileServer struct {
	file_store.UnimplementedFileTransferServer
}

func (server *FileServer) SendFile(stream file_store.FileTransfer_SendFileServer) error{
	// var fileName string
	var file *os.File

	for{
		data, err := stream.Recv()
		if err!= nil {
			if file != nil {
				file.Close()
			}
			stream.SendAndClose(&file_store.FileResponse{
				Success: false,
				Message: "error recieving chunk",
			})
			break
		}

		if data.Eof {
			if file != nil {
				file.Close()
			}
			stream.SendAndClose(&file_store.FileResponse{
				Success: true,
				Message: "file recieved successfully",
			})
			break
		}

		if file == nil {
			file, err = os.Create(data.FileName)
			if err != nil{
				stream.SendAndClose(&file_store.FileResponse{
					Success: false,
					Message: "error creating file",
				})
				break
			}
		}

		_, err = file.Write(data.Chunk)
		if err != nil{
			stream.SendAndClose(&file_store.FileResponse{
				Success: false,
				Message: "error writing to file",
			})
			break
		}
	}
	return nil
}

func main(){
	listener, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Println("Error creating listener")
	}

	grpcServer := grpc.NewServer()
	file_store.RegisterFileTransferServer(grpcServer, &FileServer{})

	log.Println("gRPC server starting at post :50051...")
	err = grpcServer.Serve(listener); if err != nil{
		log.Println("Error starting server")
	}
}

package grpc

import (
	file_store "github.com/Milindcode/file_store/proto"
)

type FileServer struct {
	file_store.UnimplementedFileTransferServer
}

// func (server *FileServer) FileStorer(stream file_store.FileTransfer_SendFileServer) error {
// 	var fileName string
// 	var file *os.File

// }

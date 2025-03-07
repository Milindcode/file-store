package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	file_store "github.com/Milindcode/file_store/grpc/proto"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func FileChunkHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket Upgrade Error:", err)
			return
		}
		defer conn.Close()

		// var file *os.File
		// var fileName string

		grpc_conn, err := grpc.NewClient("grpc:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Error connecting to grpc server: %v", err)
		}
		defer grpc_conn.Close()

		grpc_client:= file_store.NewFileTransferClient(grpc_conn)
		stream, err := grpc_client.SendFile(context.Background())
		if err != nil {
			log.Fatalf("Error connecting to stream: %v", err)
		}

		chunkCount := 1
		for {
			var temp file_store.FileData
			err := conn.ReadJSON(&temp)
			if err != nil {
				log.Println("Error getting chunk:", err)
				break
			}

			if temp.Eof{
				err = stream.Send(&temp)
				if err != nil {
					err_msg := fmt.Sprintf("Error sending chunk: %v", err)
					conn.WriteJSON(map[string]interface{}{
						"status":  false,
						"message": err_msg,
					})
				}
				break
			}

			log.Printf("Sending chunk: %d", chunkCount)
			err = stream.Send(&temp)
			if err != nil {
				err_msg := fmt.Sprintf("Error sending chunk: %v", err)
				conn.WriteJSON(map[string]interface{}{
					"status":  false,
					"message": err_msg,
				})
				break
			}
			chunkCount++
		}

		response, err := stream.CloseAndRecv()
		if err != nil {
			err_msg := fmt.Sprintf("File upload unsuccessful: %v", err)
			conn.WriteJSON(map[string]interface{}{
				"status":  false,
				"message": err_msg,
			})
		}

		conn.WriteJSON(response)

		log.Println("Closing connection...")
	}
}

package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type requestMessage struct {
	EOF      bool   `json:"eof"`
	Chunk    []byte `json:"chunk"`
	FileName string `json:"filename"`
}

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

		var file *os.File
		var fileName string

		for {
			var temp requestMessage
			err := conn.ReadJSON(&temp)
			if err != nil {
				log.Println("Error getting chunk:", err)
				if file != nil {file.Close()}
				break
			}

			if temp.EOF{
				if file != nil {
					file.Close()
				}

				conn.WriteJSON(map[string]interface{}{
					"status":  "success",
					"message": "file recieved successfully",
				})
				break
			}

			if file == nil {
				fileName = temp.FileName
				file, err = os.Create(fileName)
				if err != nil {
					conn.WriteJSON(map[string]interface{}{
						"status":  "failed",
						"message": "file cannmot be created in os",
					})
					break
				}
			}

			_, err = file.Write(temp.Chunk)
			if err != nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  "failed",
					"message": "error writing chunk",
				})
				if file != nil {file.Close()}
				break
			}
		}

		log.Println("Closing connection...")
	}
}

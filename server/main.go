package main

import (
	"log"
	"net/http"

	routes "github.com/Milindcode/file_store/server/router"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/file", routes.FileChunkHandler())


	log.Println("Starting the server at port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("Error starting the server...")
	}
}

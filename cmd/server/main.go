package main

import (
	"fmt"
	"log"

	"github.com/naolson2025/go-distributed-system/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	fmt.Println("Starting server on :8080")
	log.Fatal(srv.ListenAndServe())
}
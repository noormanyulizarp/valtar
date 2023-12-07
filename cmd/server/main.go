// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"valtar/internal/api/handlers"
)

func main() {
	handler := handlers.NewHandler()
	port := "8080"
	fmt.Printf("Server is starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

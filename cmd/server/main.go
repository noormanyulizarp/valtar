// valtar/cmd/server/main.go
package main

import (
    "fmt"
    "net/http"
    "valtar/internal/api"
)

func main() {
    handler := api.NewHandler()

    http.Handle("/", handler)

    fmt.Println("Server is starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
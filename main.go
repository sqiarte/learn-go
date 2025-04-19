package main

import (
    "fmt"
    "net/http"

    "learn-go/api" //module name + folder name
)

func main() {
    http.HandleFunc("/users", api.GetUsers)

    fmt.Println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

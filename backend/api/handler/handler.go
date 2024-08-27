package handler

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
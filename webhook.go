package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            var data map[string]interface{}
            err := json.NewDecoder(r.Body).Decode(&data)
            if err != nil {
                http.Error(w, "Bad request", http.StatusBadRequest)
                return
            }
            fmt.Printf("Received data: %+v\n", data)
            w.WriteHeader(http.StatusOK)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}

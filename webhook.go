package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Handler function that Vercel will use
func Handler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST
    if r.Method == http.MethodPost {
        var data map[string]interface{}
        
        // Decode the JSON payload
        err := json.NewDecoder(r.Body).Decode(&data)
        if err != nil {
            http.Error(w, "Bad request", http.StatusBadRequest)
            return
        }
        
        // Print the received data to the server logs
        fmt.Printf("Received data: %+v\n", data)
        
        // Respond with a 200 OK status
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Data received successfully"))
    } else {
        // Respond with a 405 Method Not Allowed for non-POST requests
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Main entry point for the Go serverless function
func main() {
    http.HandleFunc("/webhook", Handler)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}

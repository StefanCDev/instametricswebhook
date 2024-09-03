package main

import (
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello from Vercel!"))
}

func main() {
    http.HandleFunc("/webhook", Handler)
    http.ListenAndServe(":8080", nil)
}

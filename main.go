package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
		var data map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil{
			http.Error(w, "Bad request", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Received data: %+v\n", data)
		w.WriteHeader(http.StatusOK)
	}else{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}


func main() {
	http.HandleFunc("/webhook",webhookHandler)
	http.ListenAndServe(":8080", nil)
}

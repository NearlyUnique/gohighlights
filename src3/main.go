package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/view", helloJSONWorld)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to start: %v\n", err)
	}
}

func helloJSONWorld(w http.ResponseWriter, r *http.Request) {
	type message struct {
		Text string `json:"message"`
	}
	msg := message{
		Text: "Hello World",
	}

	buf, err := json.Marshal(&msg)

	if err != nil {
		log.Printf("Failed: json marshal: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error":"internal"}`))
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

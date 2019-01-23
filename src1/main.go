package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	fmt.Println(port)
	err := http.ListenAndServe(port, http.HandlerFunc(helloWorld))
	if err != nil {
		fmt.Printf("Failed to start: %v\n", err)
		os.Exit(1)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

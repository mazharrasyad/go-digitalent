package main

import (
	"fmt"
	"net/http"
)

// 1. Web Server
var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World"
	fmt.Fprint(w, msg)
}

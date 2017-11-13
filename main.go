package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer Handler")
	fmt.Fprintf(w, "Hello, "+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("0.0.0.0:9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

package main

import (
	"fmt"
	"net/http"
	//"reflect"
)

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Fuck")
	fmt.Println(request.RemoteAddr)
}

func main() {
	http.HandleFunc("/", SimpleServer)
	err := http.ListenAndServe("0.0.0.0:9007", nil)
	fmt.Println()
	if err != nil {
		panic(err)
	}
}

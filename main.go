package main

import (
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Fuck"))
}

func say(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func main() {
	http.HandleFunc("/Fuck", hello)
	http.Handle("/handle", http.HandlerFunc(say))
	http.ListenAndServe(":9001", nil)
	select {} //阻塞进程
}

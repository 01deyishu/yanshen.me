package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayFuck(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数， 默认不解析
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "Fuck ?")
}

func main() {
	http.HandleFunc("/", sayFuck)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":9007", nil)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}

package main

import (
	"fmt"
	"net"
	"net/http"
	//	"reflect"
)

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, ".")
}

func main() {
	nic, _ := net.Interfaces()
	//fmt.Println(reflect.TypeOf(nic))
	for i := 0; i < len(nic); i++ {
		addr, _ := nic[i].Addrs()
		if addr != nil {
			fmt.Printf("NIC name is : %s , Address is %s\n", nic[i].Name, addr)
		}
		//fmt.Println(reflect.TypeOf(nic[i]))
	}
	http.HandleFunc("/", SimpleServer)
	err := http.ListenAndServe("0.0.0.0:9007", nil)
	if err != nil {
		panic(err)
	}
}

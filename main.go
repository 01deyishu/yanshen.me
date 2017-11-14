package main

import (
	"fmt"
	"net"
	"net/http"
	//"reflect"
)

func SimpleServer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, ".")
}

func GetAddr() {
	nic, _ := net.Interfaces()
	for i := 0; i < len(nic); i++ {
		addr, _ := nic[i].Addrs()
		if len(addr) != 0 {
			fmt.Printf("NIC name is : %s, Address is %s\n", nic[i].Name, addr)
		}
	}
}

func main() {
	http.HandleFunc("/", SimpleServer)
	err := http.ListenAndServe("0.0.0.0:9007", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("string server ...")
	listener, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println("Error listening", err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accept ...", err.Error())
			return
		}
		fmt.Println(conn.RemoteAddr())
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Recevied date: %v", string(buf[:len]))
	}
}

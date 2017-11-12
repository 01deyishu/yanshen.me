package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	listener, err := net.Listen("tcp", "0.0.0.0:9007")
	if err != nil {
		fmt.Println("Error listenling", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting", err.Error())
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		info, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error readint", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:info]))
	}
}

package main

import (
	"fmt"
	"log"
	"net"
	"os"
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
		go writeLog(conn.RemoteAddr().String())
		go doServerStuff(conn)
	}
}

func writeLog(loginfo string) {
	logfile, err := os.OpenFile("tcp.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("logfile error")
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println(loginfo)
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		len, err := conn.Read(buf)
		tcpinfo := conn.RemoteAddr().String() + " exit"
		if err != nil {
			fmt.Println("Error reading", err.Error())
			go writeLog(tcpinfo)
			return
		}
		fmt.Printf("Recevied date: %v", string(buf[:len]))
	}
}

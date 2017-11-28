package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func writeLog(loginfo string) {
	logfile, err := os.OpenFile("tcpclient.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Create log file log")
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.Println(loginfo)
}

func main() {
	conn, err := net.Dial("tcp", "1.1.1.1:9001")
	if err != nil {
		fmt.Println("Dial 1.1.1.1:9001 error", err.Error())
		return
	}
	go writeLog(conn.LocalAddr().String())
}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func reciveTCPConn(ln *net.TCPListener) {
	for {
		err := ln.SetDeadline(time.Now().Add(time.Second * 1000))
		if err != nil {
			log.Fatal(err)
		}
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go echoHandler(conn, os.Stdout)
	}
}

func echoHandler(conn *net.TCPConn, dst io.Writer) {
	defer conn.Close()
	for {
		fmt.Println("Socket Connection")
		_, err := io.Copy(dst, conn)
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":xxxxx")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	reciveTCPConn(ln)
}

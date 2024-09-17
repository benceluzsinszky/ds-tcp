package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	fmt.Println("hello world", conn)

}

func client() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Status: ", status)

}

func server() {
	// Server
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(conn)
	}

}

func main() {

	go server()

	client()

	time.Sleep(30 * time.Second)
}

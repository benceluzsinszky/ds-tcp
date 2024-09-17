package main

import (
	"net"
	"fmt"

)
func handleConnection(conn net.Conn){
	fmt.Printf("hello world", conn)

}

func main() {
//Server
ln, err := net.Listen("tcp", ":8080")
if err != nil {
	// handle error
}
for {
	conn, err := ln.Accept()
	if err != nil {
	}
	go handleConnection(conn)
}
		
	}


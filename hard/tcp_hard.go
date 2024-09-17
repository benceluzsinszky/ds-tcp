package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
	"encoding/binary"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // Ensure the connection is closed when the function returns

	fmt.Println("Handling connection from:", conn.RemoteAddr())

}

func client() {
	conn, err := net.Dial("tcp", "localhost:8081")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when done


	// Send a message to the server
	message := 100
	//fmt.Fprintf(conn, message)
	sendBytes(conn, message)


	//fmt.Println("Message sent to server:", message)

	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return
	}
	fmt.Printf("Received response from server: %s", response)
	//message1, err := bufio.NewReader(conn).ReadString('\n')

	//fmt.Printf("Received message from server: %s", message1)
		//response := "Message received!\n"
		//_, err = conn.Write([]byte(response))

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
		message := receiveBytes(conn)// Read data into the buffer
		//message, err := bufio.NewReader(conn).ReadString('\n')
		
		
		fmt.Printf("Received message from client: %d", message)
		response := "300\n"
		_, err = conn.Write([]byte(response))

		//message1 := "Hello, Client!\n"
		//fmt.Fprintf(conn, message1)
		//fmt.Println("Message sent to client:", message1)
		defer conn.Close() // Ensure the connection is closed when done

		go handleConnection(conn)
	}



}

func sendBytes(conn net.Conn, message int){
	buf := make([]byte, 4)             // Create a buffer to hold the integer
	binary.BigEndian.PutUint32(buf, uint32(message)) // Convert integer to byte slice
	conn.Write(buf) 


}

func receiveBytes(conn net.Conn)uint32{
	responseBuf := make([]byte, 4)
	_, err := conn.Read(responseBuf)
	if err != nil {
		fmt.Println("Error reading from server:", err)
		return 0
	}
	response := binary.BigEndian.Uint32(responseBuf) // Convert byte slice to integer
	return response
}

func main() {

	go server()

	client()

	time.Sleep(30 * time.Second)
}

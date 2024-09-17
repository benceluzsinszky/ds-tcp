package main

import (
	"fmt"
	"net"
	"time"
)

type Message struct {
	seq int
	ack int
}

// sendMessage sends a Message struct over a TCP connection
func sendMessage(conn net.Conn, message Message) error {
	_, err := fmt.Fprintf(conn, "%d %d\n", message.seq, message.ack)
	return err
}

// receiveMessage returns a Message struct from a TCP connection
func receiveMessage(conn net.Conn) (Message, error) {
	var message Message
	_, err := fmt.Fscanf(conn, "%d %d\n", &message.seq, &message.ack)
	return message, err
}

func client() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close() // Ensure the connection is closed when done

	seq := 100
	message := Message{seq: seq, ack: 0}
	err = sendMessage(conn, message)
	if err != nil {
		fmt.Println("Error sending message: ", err)
		return
	}
	fmt.Println("Client sends SYN")

	message, err = receiveMessage(conn)
	if err != nil {
		fmt.Println("Error receiving message: ", err)
		return
	}
	if message.ack == seq+1 {
		fmt.Println("Client receives SYN and ACK")
		ack := message.seq + 1
		seq = seq + 1
		message = Message{seq: seq, ack: ack}
		fmt.Println("Client sends ACK")
		err = sendMessage(conn, message)
		if err != nil {
			fmt.Println("Error sending message: ", err)
			return
		}
	}
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
		defer conn.Close() // Ensure the connection is closed when done

		seq := 300

		message, err := receiveMessage(conn)
		if err != nil {
			fmt.Println("Error receiving message: ", err)
			return
		}

		client_seq := message.seq
		if message.ack == 0 {
			fmt.Println("Server receives client's SYN")
			ack := message.seq + 1
			message = Message{seq: seq, ack: ack}
			fmt.Println("Server sends SYN and ACK")
			err = sendMessage(conn, message)
			if err != nil {
				fmt.Println("Error sending message: ", err)
				return
			}
		}

		message, err = receiveMessage(conn)
		if message.ack == seq+1 && message.seq == client_seq+1 {
			fmt.Println("Server receives ACK")
			fmt.Println("Handshake complete")
		}
	}
}

func main() {
	go server()
	client()
	time.Sleep(30 * time.Second)
}

package main

import (
	"fmt"
	"time"
)

func client(syn chan int, ack chan int) {
	seq_client := 100

	syn <- seq_client
	fmt.Println("Client sends SYN")
	ack_server := <-ack
	seq_server := <-syn
	fmt.Println("Client recieves server's SYN and ACK")

	fmt.Println("Client validates server's ACK")
	if ack_server != seq_client+1 {
		fmt.Println("Invalid ACK")
		return
	}
	fmt.Println("Client sends ACK")
	ack <- seq_server + 1
	syn <- seq_client + 1
}

func server(syn chan int, ack chan int) {
	seq_server := 300

	seq_client := <-syn
	fmt.Println("Server receives client's SYN")
	syn <- seq_server
	ack <- seq_client + 1
	fmt.Println("Server sends SYN and ACK")
	ack_client := <-ack
	seq := <-syn

	fmt.Println("Server receives client's ACK")
	fmt.Println("Server validates client's ACK")
	if ack_client != seq_server+1 || seq != seq_client+1 {
		fmt.Println("Invalid ack")
		return
	}
	fmt.Println("Connection estabilished")

}

func main() {
	syn := make(chan int, 1)
	ack := make(chan int, 1)

	go client(syn, ack)
	go server(syn, ack)

	time.Sleep(30 * time.Second)
}

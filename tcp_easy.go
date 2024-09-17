package main

import (
	"fmt"
	"time"

)


func client(syn chan int, ack chan int) {
	syn_x:= 100

	//1.sends syn x
	syn  <- syn_x
	fmt.Println("Client sends syn x")

	//4. recieves syn y and ack_y
	
	ack_y := <- ack
	syn_y := <- syn
	fmt.Println("Client recieves server's syn y and ack_y")

	//5. validate ack 
	fmt.Println("Client validates server s ack")
	if(ack_y != syn_x +1){
	fmt.Println("Invalid ack")
		return
	}
	fmt.Println("Client sends ack and syn")

	//6. sends ack = y +1, sends syn x + 1
	ack <- syn_y +1
	syn <- syn_x +1

}

func server(syn chan int, ack chan int) {
	syn_y:= 300

	
	//2.receives syn x
	syn_x := <- syn
	fmt.Println("Server receives client s syn ")


	//3.sends ack = x +1 , syn = y
	syn <- syn_y

	ack <- syn_x +1
	fmt.Println("Server sends ack = x +1 , syn = y")



	//7. receive the ack from client
	ack_x := <- ack
	seq := <- syn
	fmt.Println("Server receives client s ack from client")


	//8. validate ack 
	fmt.Println("Server validates client s ack and seq")

	if(ack_x != syn_y +1 || seq != syn_x +1 ){
		fmt.Println("Invalid ack")
			return
		}

	fmt.Println("Connection estabilished")
	
	

}

func checkChannel(signal chan bool) {
	for {
		// Wait for a signal from a philosopher to request or release the fork
		signal <- true // Fork is available (this is the signal that a philosopher can take the fork)
		<-signal       // Wait for a signal from a philosopher to release the fork
	}
}





func main() {
	
	syn := make(chan int, 1)
	ack := make(chan int, 1)

	go client(syn, ack)
	go server(syn, ack)


	time.Sleep(30 * time.Second)
}

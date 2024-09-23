# ds-tcp
GoLang implementation of TCP Protocol for ITU Distributed Systems 2024

### a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

In tcp_easy we use fmt and time packages and in tcp_hard we use those as well as the net package.

In the easy implementation, we use go routines and channels to implement the 3-way handshake. The channels (chan int) were used as the data structure to transmit both data (sequence numbers) and metadata (ACKs). The *syn* channel is used to transmit the SYN (synchronize) and sequence numbers, and the *ack* channel is used to transmit ACKs (acknowledgment numbers).

In the hard implementation, we are making use of the net package's .Listen() and .Dial() functions to simulate client-server communication. These methods create real TCP connections over the network, allowing the client and server to exchange messages as if they were running on separate machines. The sendMessage and receiveMessage functions handle the transmission of sequence and acknowledgment numbers between the client and server. The server listens on a specific port (8081), while the client connects to this port. The three-way handshake (SYN, SYN-ACK, ACK) is performed through real TCP socket communication, simulating a more realistic TCP handshake process.

### b) Does your implementation use threads or processes? Why is it not realistic to use threads?

We have created a client and a server using Go routines to simulate a TCP connection. In this basic implementation, we are establishing communication between the client and server using the 3-way handshake protocol. By using channels to exchange SYN and ACK values between the client and server, we simulate how these values would be transmitted in a real TCP connection. Once the acknowledgment values are validated on both sides, the program prints "Connection established," indicating the successful completion of the handshake. These goroutines simulate the client and server, but they are not separate OS processes or threads. They are lightweight, managed by Go’s runtime.

Using threads (or Go routines in this case) to simulate a TCP connection is not realistic because the TCP protocol is designed to operate across a network. In a real TCP connection, the client and server would exchange SYN and ACK packets over the network using IP addresses and ports, with network delays, packet loss, and other factors influencing the communication.

### c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

To handle message re-ordering we could use Selective Acknowledgments (SACK) as an improvement. Unlike standard acknowledgments that only indicate the last successfully received packet, SACK allows the receiver to inform the sender about all received segments, even if some are missing. This means the sender can focus on retransmitting only the lost packets rather than resending all subsequent ones. As a result, SACK enhances network efficiency and improves throughput, particularly in environments with higher packet loss.

### d) In case messages can be delayed or lost, how does your implementation handle message loss?

Currently, we do not handle delays or lost messages. 

If messages are delayed, we would need to implement a timeout mechanism. A timeout ensures that the sender waits for an acknowledgment (ACK) for a certain period of time. If the ACK is not received within the specified timeout window, the sender assumes that the packet or the acknowledgment was lost. If a message (such as a SYN, SYN-ACK, or ACK) is not received within the timeout period, the sender could retransmit the message, assuming it was lost or delayed beyond the acceptable window.

In the case of message loss, Selective Acknowledgment (SACK) would be an efficient way to handle the retransmission of only the lost packets. The receiver would notify the sender of which packets were successfully received, allowing the sender to retransmit only the lost ones.

### e) Why is the 3-way handshake important?
The 3-way handshake is crucial for establishing a reliable connection between the sender and receiver. It ensures that both sides exchange and synchronize sequence numbers and acknowledgments, which are incremented by one when received, allowing each end to validate the communication. If the sequence numbers don’t follow the expected order, the receiver can detect and flag them as incorrect, preventing errors and ensuring that data transmission is properly synchronized.

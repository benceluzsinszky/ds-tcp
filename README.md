# ds-tcp
GoLang implementation of TCP Protocol for ITU Distributed Systems 2024
a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
So for tcp_easy we are using fmt and time packages, wherease for tcp_hard we used also package called net.

We are using net package to implement the 3-handshake.
The channels were used (chan int) as the data structure to transmit both data (sequence numbers) and metadata (ACKs).

- The syn channel is used to transmit the SYN (synchronize) and sequence numbers.
- The ack channel is used to transmit ACKs (acknowledgment numbers).

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

Our implementation uses go client(syn, ack) and go server(syn, ack), that is creating two concurrent goroutines.

These goroutines simulate the client and server, but they are not separate OS processes or threads. They are lightweight, managed by Go’s runtime.
So it is not realistic to use neither the threads nor the proccesses as they are hardwer dependent, wherease a goroutines are managed by Go runtime.
If we were to run our codebase on distinct servers, we would use the processes that will be triggered once a port has been accessed.

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

d) In case messages can be delayed or lost, how does your implementation handle message loss?
If messages are delayed, we would need to implement a timeout mechanism. In the real world, a timeout ensures that the sender waits for an acknowledgment (ACK) for a certain period. If the ACK is not received within the specified timeout window, the sender assumes that the packet or the acknowledgment was lost.
In our Go implementation, this would mean introducing a timer. If a message (such as a SYN, SYN-ACK, or ACK) is not received within the timeout period, the sender could retransmit the message, assuming it was lost or delayed beyond the acceptable window.

In the case of message loss, Selective Acknowledgment (SACK) would be an efficient way to handle the retransmission of only the lost packets. The receiver would notify the sender of which packets were successfully received, allowing the sender to retransmit only the lost ones.

e) Why is the 3-way handshake important?
The 3-way handshake is crucial for establishing a reliable connection between the sender and receiver. It ensures that both sides exchange and synchronize sequence numbers and acknowledgments, which are incremented by one when received, allowing each end to validate the communication. If the sequence numbers don’t follow the expected order, the receiver can detect and flag them as incorrect, preventing errors and ensuring that data transmission is properly synchronized.

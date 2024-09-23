# ds-tcp
GoLang implementation of TCP Protocol for ITU Distributed Systems 2024
a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?
So for tcp_easy we are using fmt and time packages, wherease for tcp_hard we used also package called net.

We are using net package to implement the 3-handshake.
The channels were used (chan int) as the data structure to transmit both data (sequence numbers) and metadata (ACKs).

- The syn channel is used to transmit the SYN (synchronize) and sequence numbers.
- The ack channel is used to transmit ACKs (acknowledgment numbers).

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

Our implementation doesn't use OS-level threads rather goroutines, which are more efficient and scalable than threads for concurrency in Go, especially for high-concurrency applications.


c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?

d) In case messages can be delayed or lost, how does your implementation handle message loss?

e) Why is the 3-way handshake important?

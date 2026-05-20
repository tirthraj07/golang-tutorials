package main

import "fmt"

/*
Buffered Channels in Go

So far we have seen unbuffered channels -> where the sender and receiver must MEET at the same time

Buffered Channels change this -
They introduce temporary storage, meaning -> sender may continue without immediate receiver

Syntax
ch := make(chan type, capacity=1)
Where capacity is the size of the "Queue"

So go will not block the sender from sending the messages until we hit the capacity
RULES:
	- Sender is blocked IF buffer is full
	- Receiver is blocked IF buffer is empty

This is a huge difference as in unbuffered channels as send waits for receiver immediately

NOTE:
Buffered Channels Are Queues - First In First Out
*/

func main() {
	buffer_size := 5
	ch := make(chan int, buffer_size)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	// Note: Here ch <- 6 would have blocked the channel

	close(ch) // Remember to close the channel

	fmt.Println("All values sent!")

	for val := range ch {
		fmt.Println(val)
	}

	/*
		IF YOU DON'T CLOSE THE CHANNEL -
		for val := range ch {
			fmt.Println(val)
		}
		This threw an error: all goroutines are asleep - deadlock!
		WHY? range ch means: “Keep receiving values UNTIL the channel is CLOSED.”
		After the buffer was empty, the loop thinks - "More values may still come later"
		Thus it waits for next receive.
		Receive Blocks

		Because:
			buffer empty
			no sender exists

		Meanwhile Main Goroutine Also Blocked - No other runnable goroutines remain.

		Mental Model shift - Channels are streams, not collections
		A channel is NOT merely queue of values
		It represents: an ongoing communication stream
		And closing signals: end-of-stream
	*/

}

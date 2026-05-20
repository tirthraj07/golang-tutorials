package main

import (
	"fmt"
	"time"
)

/*
Channels
Channels are how goroutines communicate with each other
“Do not communicate by sharing memory; share memory by communicating.”

What Is a Channel?

A channel is: a typed communication pipe

Used to:
	send data
	receive data
	synchronize goroutines


*/

func consumer(ch chan string) {
	// Process Strings here
	/*
		value := <-ch
		BLOCKS until: value available
	*/
	for str := range ch {
		fmt.Println("[CONSUMER] Processing String : ", str)
		time.Sleep(500 * time.Millisecond)
	}
}

func produceString(str string, ch chan string) {
	fmt.Println("[PRODUCER] Sending String : ", str)
	// SENDING DATA : data flows INTO channel
	/*
		BLOCKS until: receiver ready
	*/
	ch <- str
}
func main() {
	// Creating channels
	ch := make(chan string)
	go consumer(ch)

	produceString("Hello", ch)
	produceString("World", ch)
	produceString("!", ch)
	// sender should close channel
	close(ch)
}

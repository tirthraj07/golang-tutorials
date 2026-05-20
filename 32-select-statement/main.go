package main

import (
	"fmt"
	"time"
)

/*
Select Statement in Go Language

In Go, the select statement allows you to wait on multiple channel operations, such as sending or receiving values.
Similar to a switch statement, select enables you to proceed with whichever channel case is ready, making it ideal for handling asynchronous tasks efficiently.

Consider a scenario where two tasks complete at different times.
We’ll use select to receive data from whichever task finishes first.
*/

func task1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Task 1 Completed"
}

func task2(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "Task 2 Completed"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go task1(ch1)
	go task2(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	close(ch1)
	close(ch2)

	/*
		After 2 seconds, Task 1 completed will be printed because task1 completes before task2.
		If task2 finished first, then Task 2 completed would be printed.


		Syntax:
		select {
			case value := <-channel1:
				// Executes if channel1 is ready to send/receive
			case channel2 <- value:
				// Executes if channel2 is ready to send/receive
			default:
				// Executes if no other case is ready
		}

		select waits until at least one channel operation is ready.
		If multiple cases are ready, one is chosen at random.
		The default case executes if no other case is ready, avoiding a block.
	*/

	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No channels are ready")
	}

	/*
		Explanation: Here, the default case ensures that select doesn’t block if no channels are ready, printing "No channels are ready".
	*/

	// Infinite Blocking without Cases
	// If a select statement is empty (i.e., contains no cases), it blocks indefinitely. This is often used in cases where an infinite wait is necessary, but here’s what it looks like using our channels.
	// select {}  // This blocks forever as there are no cases
	// Explanation: Since there are no cases, select will block permanently, causing a deadlock if there are no other goroutines active.

}

package main

import (
	"fmt"
	"time"
)

// You can use channels for synchronizing the threads as well - goroutine synchronization using channels without using WaitGroup.
/*
This Pattern Is Often Called
	completion signaling
	or:
	fan-in synchronization
*/
func processing(ch chan struct{}, workerID int) {
	defer func() {
		ch <- struct{}{}
	}()
	fmt.Printf("[%v] Processing...\n", workerID)
	time.Sleep(time.Second * 3)
	fmt.Printf("[%v] Finished\n", workerID)
}

func main() {
	ch := make(chan struct{})

	workerCount := 10

	for workerID := range workerCount {
		go processing(ch, workerID)
	}

	for range workerCount {
		<-ch // Blocks until all the workers have not signalled done
	}

	fmt.Println("Main Execution Finished")
}

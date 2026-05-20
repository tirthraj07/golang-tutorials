package main

import (
	"fmt"
	"sync"
	"time"
)

// Here we are using 'done' channel as completion signaling. We have already seen this in tut-30: Completion Signaling
func worker(workerID int, messages chan string, done chan struct{}, wg *sync.WaitGroup) {
	// Guaranteed Cleanup
	defer func() {
		fmt.Printf("[%v] Cleaning up..\n", workerID)
		time.Sleep(2 * time.Second)
		wg.Done()
	}()

	fmt.Printf("[%v] Ready..\n", workerID)
	for {
		select {
		case msg, ok := <-messages: // Add the 'ok' boolean
			if !ok {
				// If the channel is closed, set it to nil.
				// This disables this specific 'case', preventing the infinite loop.
				messages = nil
				continue
			}
			time.Sleep(2 * time.Second)
			fmt.Printf("[%v] %v\n", workerID, msg)
		case <-done:
			return
		}
	}
}

func main() {
	messages := make(chan string)
	done := make(chan struct{})
	var wg sync.WaitGroup
	numWorkers := 5

	for workerID := range numWorkers {
		wg.Add(1)
		go worker(workerID, messages, done, &wg)
	}

	for i := range 10 {
		messages <- fmt.Sprintf("Message %d", i)
	}

	close(messages)

	close(done) // Calling close(done) immediately after broadcasts the signal to all workers. This is called Broadcast Cancellation
	wg.Wait()
	fmt.Println("Main Thread Finished Execution")
}

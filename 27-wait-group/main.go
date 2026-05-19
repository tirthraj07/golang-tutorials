package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func worker(worker_id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range 10 {
		fmt.Printf("[THREAD %v] %v\n", worker_id, i)
		time.Sleep(500 * time.Millisecond)
	}

}

func main() {
	parallelRun := false
	args := os.Args
	if len(args) == 2 && args[1] == "--parallel" {
		parallelRun = true
	}

	var wg sync.WaitGroup

	startTime := time.Now()
	for worker_id := range 10 {
		wg.Add(1)
		if parallelRun {
			go worker(worker_id, &wg)
		} else {
			worker(worker_id, &wg)
		}
	}

	wg.Wait()
	endTime := time.Now()
	fmt.Println("Main Finished Execution")
	fmt.Println("Total Time : ", endTime.Sub(startTime))
}

/*
Sequential : 50.0648923s
Parallel   : 5.0090511s
*/

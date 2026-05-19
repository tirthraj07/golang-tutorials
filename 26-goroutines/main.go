package main

import (
	"fmt"
	"time"
)

/*
Goroutines
They are lightweight concurrent functions
They allow Go programs to:
	handle massive concurrency
	efficiently utilize CPUs
	scale backend systems elegantly

First Important Concept
Concurrency != Parallelism

Concurrency -> Managing multiple tasks: seemingly at same time
Parallelism -> Actually executing simultaneously on: multiple CPU cores

Go Supports BOTH
Goroutines enable:
	concurrency naturally
	parallelism when CPUs available

A goroutine is a lightweight managed thread created using: `go`

HUGE Difference From OS Threads Goroutines are: NOT 1:1 OS threads. Go runtime manages them intelligently
Goroutines Are Tiny

Initial stack:
~2 KB

Threads often:
~1 MB

Huge difference.

Goroutines are: nondeterministic. Scheduler decides execution order.
*/

func worker(name int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[THREAD %v] %v\n", name, i)

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	for worker_id := range 10 {
		go worker(worker_id)
	}

	time.Sleep(10 * time.Second)
	fmt.Println("Main Execution Finished")
}

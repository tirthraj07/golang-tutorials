package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
What problem are we trying to solve?
Suppose we have counter++
Internally it is 3 separate operations -> READ, MODIFY, WRITE (all of these are not atomic)

Traditional Solution - use Mutex
```
mu.Lock()
counter++
mu.Unlock()
```
Correct.
But:

locking overhead
scheduler involvement
contention cost


Atomics Provide
hardware-level synchronization
without full mutex locking.

Go Atomic Package - sync/atomic
Important Methods
Basic API 		 - var counter atomic.Int64
Atomic Increment - counter.Add(1)
Atomic Read 	 - counter.Load()
Atomic Store     - counter.Store(42)


CAS API - Compare and Swap
CompareAndSwap(old, new)
Returns:
	true if swap succeeded
	false otherwise

CAS enables: lock-free algorithms
Core primitive behind:
	concurrent queues
	lock-free stacks
	schedulers
	runtime internals



*/

func main() {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			counter.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println(counter.Load())

	var value atomic.Int64
	value.Store(10)
	swapped := value.CompareAndSwap(10, 20)
	fmt.Println(swapped)
	fmt.Println(value.Load())
}

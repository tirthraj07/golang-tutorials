package main

import (
	"fmt"
	"sync"
	"time"
)

/*
In go, sync.Mutex treats both read and writes the same
Meaning - only ONE goroutine allowed at a time
even if: everybody only reading

That can severely hurt throughput in:
	caches
	config stores
	metadata systems
	routing tables
	read-heavy systems

RWMutex Solves This
RWMutex distinguishes:
	readers
	vs
	writers

Multiple Readers Allowed Simultaneously - R R R R R allowed
Writer Requires Exclusive Access - W
Exclusive lock blocks all readers and all writers

Basic API var mu sync.RWMutex
Read Lock 		mu.RLock()
Read Unlock 	mu.RUnlock()
Write Lock 		mu.Lock()
Write Unlock 	mu.Unlock()
*/

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.value
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup

	// Readers
	for i := range 5 {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			for range 5 {
				fmt.Printf(
					"[Reader %d] Value = %d\n",
					id,
					counter.Get(),
				)

				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	// Writer
	wg.Add(1)

	go func() {
		defer wg.Done()

		for range 5 {
			counter.Increment()

			fmt.Println("[Writer] Incremented")

			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}

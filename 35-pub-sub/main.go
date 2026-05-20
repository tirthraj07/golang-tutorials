package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Core problem - msg <- ch consumes message. But only one receiver gets it
If you share the channel with multiple consumers, only one consumer will receive the message

But pub-sub requires
Publisher -> ALL subscribers receive copy
i.e fan-out duplication

You must:
maintain subscriber list
broadcast message to each subscriber channel
*/

type Subscriber[T any] struct {
	id         int
	queue      chan T
	numMessage int
}

type PubSub[T any] struct {
	subscribers []*Subscriber[T]

	mu sync.RWMutex
}

func (ps *PubSub[T]) Subscribe(
	id int,
	bufferSize int,
	handler func(T),
	wg *sync.WaitGroup,
) {
	sub := &Subscriber[T]{
		id:    id,
		queue: make(chan T, bufferSize),
	}

	ps.mu.Lock()
	ps.subscribers = append(ps.subscribers, sub)
	ps.mu.Unlock()

	// Dedicated worker for subscriber
	go func() {
		defer wg.Done()
		for msg := range sub.queue {
			sub.numMessage += 1
			handler(msg)
		}
	}()
}

func (ps *PubSub[T]) Publish(msg T) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, sub := range ps.subscribers {

		// Non-blocking enqueue
		select {
		case sub.queue <- msg:

			// default:
			// 	fmt.Printf(
			// 		"[Subscriber %d] Queue Full -> Dropping Message\n",
			// 		sub.id,
			// 	)
			// This is handling the pressure on the consumer side by dropping messages
			// Commenting this is handling the pressure on producer side where we delay the producer from sending the message to the channel
		}
	}
}

func (p *PubSub[T]) CloseChannels() {
	for _, sub := range p.subscribers {
		close(sub.queue)
		fmt.Printf("[Subscriber %v] Channel Closed\n", sub.id)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ps := PubSub[string]{}

	ps.Subscribe(1, 5, func(msg string) {
		time.Sleep(1 * time.Second)
		fmt.Println("[Subscriber 1]", msg)
	}, &wg)

	ps.Subscribe(2, 5, func(msg string) {
		time.Sleep(2 * time.Second)
		fmt.Println("[Subscriber 2]", msg)
	}, &wg)

	ps.Subscribe(3, 5, func(msg string) {
		time.Sleep(5 * time.Second)
		fmt.Println("[Subscriber 3]", msg)
	}, &wg)

	for i := range 20 {
		ps.Publish(fmt.Sprintf("Message %d", i))
	}

	ps.CloseChannels() // signal goroutines to stop
	wg.Wait()          // now they can actually finish and call Done()

	// Safe to read numMessage — all goroutines have exited
	for _, sub := range ps.subscribers {
		fmt.Printf("[Subscriber %v]Total Messages %v\n", sub.id, sub.numMessage)
	}
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Sequence int
	Value    string
}

type EventLog struct {
	buffer   []Message
	capacity int

	start int // oldest message sequence
	end   int // next sequence number

	mu sync.RWMutex
}

func NewEventLog(capacity int) *EventLog {
	return &EventLog{
		buffer:   make([]Message, capacity),
		capacity: capacity,
	}
}

func (el *EventLog) Publish(value string) {
	el.mu.Lock()
	defer el.mu.Unlock()

	msg := Message{
		Sequence: el.end,
		Value:    value,
	}

	idx := el.end % el.capacity

	el.buffer[idx] = msg

	el.end++

	// Buffer overflow -> move start forward
	if el.end-el.start > el.capacity {
		el.start++
	}
}

func (el *EventLog) Read(sequence int) (Message, bool) {
	el.mu.RLock()
	defer el.mu.RUnlock()

	// Consumer too slow -> message overwritten
	if sequence < el.start {
		return Message{}, false
	}

	// Message not produced yet
	if sequence >= el.end {
		return Message{}, false
	}

	idx := sequence % el.capacity

	return el.buffer[idx], true
}

func (el *EventLog) CurrentStart() int {
	el.mu.RLock()
	defer el.mu.RUnlock()

	return el.start
}

func (el *EventLog) CurrentEnd() int {
	el.mu.RLock()
	defer el.mu.RUnlock()

	return el.end
}

type Subscriber struct {
	id     int
	offset int
}

func NewSubscriber(id int, startOffset int) *Subscriber {
	return &Subscriber{
		id:     id,
		offset: startOffset,
	}
}

func (s *Subscriber) Consume(
	log *EventLog,
	processTime time.Duration,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		msg, ok := log.Read(s.offset)

		if !ok {

			// Detect if consumer lagged behind retention
			currentStart := log.CurrentStart()

			if s.offset < currentStart {
				fmt.Printf(
					"[Subscriber %d] Missed messages. Skipping from %d -> %d\n",
					s.id,
					s.offset,
					currentStart,
				)

				s.offset = currentStart
			}

			// Wait for new messages
			time.Sleep(100 * time.Millisecond)

			// Exit condition for demo
			if s.offset >= 20 {
				return
			}

			continue
		}

		fmt.Printf(
			"[Subscriber %d] Consumed: %s\n",
			s.id,
			msg.Value,
		)

		s.offset++

		time.Sleep(processTime)
	}
}

func main() {
	log := NewEventLog(5)

	var wg sync.WaitGroup

	sub1 := NewSubscriber(1, 0)
	sub2 := NewSubscriber(2, 0)
	sub3 := NewSubscriber(3, 0)

	wg.Add(3)

	go sub1.Consume(log, 1*time.Second, &wg)
	go sub2.Consume(log, 2*time.Second, &wg)
	go sub3.Consume(log, 4*time.Second, &wg)

	// Publisher
	for i := range 20 {
		log.Publish(fmt.Sprintf("Message %d", i))

		fmt.Printf("[Publisher] Published Message %d\n", i)

		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()
}

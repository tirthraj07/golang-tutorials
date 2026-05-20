package main

import (
	"fmt"
	"os"
	"slices"
	"sync"
	"time"
)

type Post struct {
	PostID   string
	Username string
	ImageURL string
	Likes    int

	mu sync.Mutex
}

func (p *Post) incrementLikes(workerID int, wg *sync.WaitGroup, useMutex bool) {
	defer func() {
		if useMutex {
			p.mu.Unlock()
		}
		wg.Done()
	}()
	fmt.Printf("[%v] Incrementing Post Like\n", workerID)
	time.Sleep(3 * time.Second)
	if useMutex {
		p.mu.Lock()
	}
	p.Likes += 1

}

func main() {
	useMutex := false
	if slices.Contains(os.Args, "--use-mutex") {
		useMutex = true
	}

	myPost := Post{
		PostID:   "123",
		Username: "tirthraj07",
		ImageURL: "s3://<my-bucket>/images/123",
		Likes:    0,
	}
	var wg sync.WaitGroup
	numWorkers := 200000
	startTime := time.Now()
	for workerID := range numWorkers {
		wg.Add(1)
		go myPost.incrementLikes(workerID, &wg, useMutex)
	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println("-- Finished --")
	fmt.Printf("Expected %v\nGot %v\nDiff %v\nTime %v", numWorkers, myPost.Likes, numWorkers-myPost.Likes, endTime.Sub(startTime))
}

/*
Without Mutex :
-- Finished --
Expected 200000
Got 199190
Diff 810
Time 12.1593857s

With Mutex
-- Finished --
Expected 200000
Got 200000
Diff 0
Time 13.3033596s
*/

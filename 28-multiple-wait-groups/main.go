/*
Problem Statement: Can you have multiple wait groups?
Say I ran 5 threads in group 1 and 5 threads in group 2.
If group 1 finishes -> I want to print "Finished with Group 1" and
similarly if group 2 finishes, i want to print "Finished with Group 2"
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(group string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Duration(id) * time.Second)

	fmt.Printf("[%s] Worker %d done\n", group, id)
}

func main() {
	var group1 sync.WaitGroup
	var group2 sync.WaitGroup

	var final sync.WaitGroup

	final.Add(2)

	for i := 1; i <= 5; i++ {
		group1.Add(1)

		go worker("Group 1", i, &group1)
	}

	for i := 1; i <= 5; i++ {
		group2.Add(1)

		go worker("Group 2", i, &group2)
	}

	go func() {
		defer final.Done()

		group1.Wait()

		fmt.Println("Finished with Group 1")
	}()

	go func() {
		defer final.Done()

		group2.Wait()

		fmt.Println("Finished with Group 2")
	}()

	final.Wait()
}

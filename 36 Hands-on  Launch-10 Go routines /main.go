package main

import (
	"fmt"
	"sync"
)

func main() {

	channels := make(chan int)

	var wg sync.WaitGroup

	// Launch 10 Go Routines

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				channels <- index*10 + j
			}

		}(i)
	}

	// Close Channel when all
	go func() {
		wg.Wait()
		close(channels)

	}()
	// Read From Channel
	for n := range channels {
		fmt.Println("Stations", n)
	}

	fmt.Println("All Channels are closed")
}

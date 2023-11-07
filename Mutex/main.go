package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// We use Mutex to prevents multiple goroutines from accessing a variable at the same time to prevent race conditon.
	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock() // Here, we Lock the variable using muetx
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mtx.Unlock() // Here, the variable gets unlocked.
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumCPU())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumCPU())
	fmt.Println("Count:", counter)
}

// i do not understand this code totally.
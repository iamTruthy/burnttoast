package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("CPUs:", runtime.NumCPU()) // This prints the number of CPUs
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // This prints the number of Go routines

	go func() {
		fmt.Println("Hello from Alice")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from The Mad Hatter")
		wg.Done()
	}()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}

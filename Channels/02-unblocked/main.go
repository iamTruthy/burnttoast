package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Goroutines", runtime.NumGoroutine())

	// Using a func literal to unblock a channel
	ch := make(chan int)

	go func() {

		ch <- 42

	}()
	fmt.Println("Goroutines", runtime.NumGoroutine()) // There are two go routines here.

	fmt.Println(<-ch)
	fmt.Println("Goroutines", runtime.NumGoroutine())

}



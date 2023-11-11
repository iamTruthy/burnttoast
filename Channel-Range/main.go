package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Goroutine", runtime.NumGoroutine()) // to see number(s) of Goroutines running
	fmt.Println("CPUs", runtime.NumCPU()) // To see the nubers of CPUs running (i think?)
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)

	}()

	fmt.Println("Goroutine", runtime.NumGoroutine())

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Goroutine", runtime.NumGoroutine())

	fmt.Println("Exiting now...")
	fmt.Println("Goroutine", runtime.NumGoroutine())

}

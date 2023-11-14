package main

import "fmt"

// using a func literal to unblock a channel

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

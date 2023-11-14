package main

import "fmt"

// Using Buffered Channel to unblock a channe

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case v := <-q:
			fmt.Println("From channel q: ", v)
		case v := <-c:
			fmt.Println("From channel c:", v)
			return
		}
	}
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)

	}()

	return c
}

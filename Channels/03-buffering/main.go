package main

import "fmt"

func main() {
	ch := make(chan int, 3) // This is BUFFERING.

	ch <- 12
	ch <- 16
	ch <- 6

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Printf("%T\n", ch)

}

/*

This code will not run because, there can only be 2 channels "(chan int, 2)"

func main() {
	ch := make(chan int,2 )

	ch <- 12
	ch <- 16
	ch <- 6

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}


*/

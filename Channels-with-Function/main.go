package main

import "fmt"

func main() {

	ch := make(chan int)

	go foo(ch)

	bar(ch)

	fmt.Println("Exiting now...")

}

func foo(ch chan<- int) {

	ch <- 33

}

func bar(ch <-chan int) {

	fmt.Println(<-ch)
}

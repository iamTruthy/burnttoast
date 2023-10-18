package main

import "fmt"

func main() {
	i := foo(1, 2, 3)
	fmt.Println(i)
}

func foo(n ...int) int {

	x := 0
	for _, v := range n {
		x += v
	}
	return x
}

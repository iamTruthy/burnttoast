package main

import "fmt"

func main() {
	var n = 53
	fmt.Println(n)

	i := &n
	fmt.Println(i)

	*i = 21
	fmt.Println(*i)

}
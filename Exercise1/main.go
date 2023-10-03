package main

import "fmt"

var num float32 = 3.33

const (
	_ = iota
	a
	b
	c
)

func main() {
	d := 0

	fmt.Printf("%v \t %b \n", b<<1, b<<1)
	fmt.Printf("%v \t %b \n", c<<2, c<<2)
	fmt.Println(d)

	fmt.Println(num)
	fmt.Println(a + num)

	fmt.Println("Hi, this from the Nanoverse. Can you see this?")
	
	/*
		this was written through the termial using the "nano main.go" command
		on the terminal
	*/
}

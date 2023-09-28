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

}

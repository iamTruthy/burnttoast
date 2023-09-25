package main

import (
	"fmt"
	"github.com/iamTruthy/Waffle"
)

func main() {
	a := waffle.Bill()
	b := waffle.Order()
	c := waffle.Ready()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

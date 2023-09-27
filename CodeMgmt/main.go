package main

import (
	"fmt"
	"github.com/iamTruthy/Waffle"
)

func main() {

	a := waffle.Order()
	fmt.Println(a)

	b := waffle.Meal()
	fmt.Println(b)

	c := waffle.Bill()
	fmt.Println(c)

}

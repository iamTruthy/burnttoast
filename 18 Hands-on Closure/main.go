package main

import (
	"fmt"
	"math"
)

func main() {
	x := increase(3)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}

func increase(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a,c)
	}
}

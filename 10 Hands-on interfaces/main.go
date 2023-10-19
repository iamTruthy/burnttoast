package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {

	s := square{
		length: 4,
		width:  4,
	}
	info(s)
	fmt.Println(info(s))

	c := circle{
		radius: 6,
	}
	info(c)
	fmt.Println(info(c))

}

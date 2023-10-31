package main

import "fmt"

func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}

// Implementing type constraint & type set with an interface
type math interface {
	~int | float64
}

// This here, is a type constraint.
func addIF[T math](a, b T) T {
	return a + b
}

// here, is an underlying type (an Alias)
type openSesami int

func main() {

	var n openSesami = 32

	fmt.Println(addFloat(6.5, 9.3))
	fmt.Println(addInt(6, 9))

	fmt.Println(addIF(3.4, 7.3))
	fmt.Println(addIF(n, 2))

}

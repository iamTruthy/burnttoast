package main

import "fmt"

func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}


//This here, is a type constraint.
func addIF[T int | float64](a, b T) T {
	return a + b

}

func main() {

	fmt.Println(addFloat(6.5, 9.3))
	fmt.Println(addInt(6, 9))

	fmt.Println(addIF(3.4, 7.3))
	fmt.Println(addIF(6, 2))

}

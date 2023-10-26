package main

import "fmt"

func main() {

	i := soln(3, 7, multiply)
	fmt.Println(i)

	x := multiply(2, 9)
	fmt.Println(x)

	fmt.Println(callSquare(square, 6))

}

func soln(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func multiply(a int, b int) int {
	return a * b
}

func callSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the square of %v is %v", a, x)
}

func square(n int) int {
	return n * n
}

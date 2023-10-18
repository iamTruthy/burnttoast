package main

import "fmt"

func main() {

	x := add(3, 4)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", doMath)

	fmt.Println("----------------- For y and sub()---------------------------")

	y := sub(34, 16)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)

	fmt.Println("------------------ For i and doMath() --------------------------")

	i := doMath(37, 64, add)
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", doMath)

	fmt.Println("------------------ For v --------------------------")

	v := doMath(76, 24, sub)
	fmt.Println(v)
	fmt.Printf("%T\n", v)

}

// This is a Callback func
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

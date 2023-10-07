package main

import "fmt"

func main() {
	a := []string{"sunhine", "bumblebee", "icecream", "vanillaice"}
	b := a
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("------------------")

	a[0] = "banana milkshake"
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("------------------")

	i := []int{2, 4, 6, 8, 16, 14}
	x := make([]int, 6)
	copy(x, i)

	fmt.Println(i)
	fmt.Println(x)

	fmt.Println("------------------")

	i[0] = 16
	fmt.Println(i)
	fmt.Println(x)

}

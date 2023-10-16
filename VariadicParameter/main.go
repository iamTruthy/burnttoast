package main

import "fmt"

func main() {
	i := num(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(i)

}

func num(n ...int) int {
	fmt.Println(n)
	fmt.Printf("%T\n",n)

	x := 0
	for _, v := range n {
		x += v
	}
	return x
}
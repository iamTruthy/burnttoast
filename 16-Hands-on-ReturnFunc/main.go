package main

import "fmt"

func main() {
	x := num()
	fmt.Println(x())

}

func num() func() int {
	return func() int {
		return 42
	}
}

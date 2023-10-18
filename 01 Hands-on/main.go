package main

import "fmt"

func main() {
	foo()
	fmt.Println(foo())

	bar()
	fmt.Println(bar())
}

func foo() int {
	return 24
}

func bar() (int, string) {
	return 21, "Savage"

}

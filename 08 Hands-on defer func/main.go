package main

import "fmt"

type animal string

func main() {
	x := foo(animal("Bear"))
	defer fmt.Println(x)

	bar()
	fmt.Println(bar())

	func() {
		defer fmt.Println("Hey Anon, i am a funcion but just incognito")
	}()

}

func foo(a animal) string {
	return "Bear"
}

func bar() int {
	return 40
}

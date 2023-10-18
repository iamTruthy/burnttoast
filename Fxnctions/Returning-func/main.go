package main

import "fmt"

func main() {
	x := foo()
	defer fmt.Println(x)

	y := bar()
	fmt.Println(y())
}

func foo() string {
	return "He is 42yrs old"
}

// This is a func bar() that returns an anonymous func(), that returns a string. 

func bar() func() string{
	return func() string {
		return "Dushane is a man"
	}
}
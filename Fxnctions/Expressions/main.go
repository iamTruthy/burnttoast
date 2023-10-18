package main

import "fmt"

func main() {

	//This is an Anonymous Func

	func() {
		fmt.Println("Hey there, i am an anonymous func")
	}()

	func(s string) {
		fmt.Println("Hi there i am also an Anon func called ", s)
	}("foo")

	// This is a Func Expression. A func expresssion is basically assigning a function to a variable.
	x := func() {
		fmt.Println("i am bar")
	}

	x()

	i := func(s string) {
		fmt.Println("My name is ", s)
	}

	i("foo-bar")

}

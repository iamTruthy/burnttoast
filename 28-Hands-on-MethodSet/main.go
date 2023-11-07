package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello World")

}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()

}

func main() {

	// This code doesn't run

	// p := person{
		// name: "Todd",
	// }
	// saySomething(p)

	ap := person{
		name: "Todd",
	}
	saySomething(&ap)

}

/*

Hands-on exercise #2
This exercise will reinforces the understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething


*/

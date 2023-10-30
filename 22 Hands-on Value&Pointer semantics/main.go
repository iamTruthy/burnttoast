package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "Wendy",
	}
	fmt.Println(p)

	p = changeName(p, "King")
	fmt.Println(p)

	changeToPointer(&p, "Donald")
	fmt.Println(p)

}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeToPointer(p *person, s string) {
	p.first = s
}

/*

Create two functions to change a field in a struct called "first" of TYPE string:

● One function will use VALUE SEMANTICS

○ this function will return some VALUE of some TYPE

● The other function will use POINTER SEMANTICS

○ this function will return nothing

● don't use methods

*/

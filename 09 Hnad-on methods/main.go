package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, " i am ", p.age, " years old")
}

func main() {

	h := person{
		first: "Todd",
		age:   37,
	}

	h.speak()

}

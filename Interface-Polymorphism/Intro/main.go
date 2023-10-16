package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p person) speak() {
	fmt.Println("i am ", p.first, p.last)
}

type sensei struct {
	person
	ninja bool
}

func (n sensei) speak() {
	fmt.Println("i am ", n.first, n.last)
}

type human interface {
	speak()
}

func sayHello(h human) {
	h.speak()
}

func main() {

	p := person{
		first: "Itadori",
		last:  "Yuji",
	}
	sayHello(p)

	n := sensei{
		person: person{
			first: "Yagamami",
			last:  "Litecoin",
		},
		ninja: true,
	}

	sayHello(n)

}

package main

import "fmt"

func main() {
	i := 42
	p := &i

	fmt.Println(i)

	fmt.Printf("i is of type %T\n", i)

	fmt.Println(p) // 'p' is a pointer to 'i' Prints out the address where "i" is stored.

	fmt.Printf("p is of type %T\n ", p)

	*p = 21         // This is also "*&i = 21" because p = &i and *p = *&i
	fmt.Println(*p) //	Dereferencing 'p' gives you the integer that 'p' points to
	fmt.Printf("*p is of type %T\n", *p)

	/*
		Dereferencing 'p' gives you the integer that 'p' points to (Gives you the value stored in that address)
		here,a new value has been assigned to "the pointer to "i" " therefore, the value if "i" changes from 42 to 21
	*/

	// You can change the value that 'p' points to
	fmt.Println(i) // The value of 'i' is now 2

}

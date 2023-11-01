package main

import (
	"fmt"
	"sort"
)

type audit struct {
	person []string
	age    []int
}

func main() {

	aud := audit{
		person: []string{"Frank", "Nate", "Marc", "Jordan", "Finn", "Todd"},
		age:    []int{9, 2, 4, 1, 4, 5, 3, 8, 15, 8, 19, 21, 4, 14},
	}
	fmt.Println(aud)

	fmt.Println("---------------------------")

	fmt.Println(aud.person)
	sort.Strings(aud.person)
	fmt.Println(aud.person)

	fmt.Println("---------------------------")

	fmt.Println(aud.age)
	sort.Ints(aud.age)
	fmt.Println(aud.age)

}

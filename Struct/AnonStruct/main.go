package main

import "fmt"

type brand struct {
	clothing string
	footware string
	watches  string
}

func main() {

	// This is an Anonymous struct

	bd := struct {
		month string
		date  int
		day   string
		year  int
	}{
		day:   "Thursday",
		month: "May",
		date:  25,
		year:  2000,
	}

	fmt.Printf("%T\n", bd)

	// This is a struct

	b := brand{
		clothing: "golf le fleur",
		footware: "Nike",
		watches:  "Lacoste",
	}

	fmt.Printf("%T\n", b)

}

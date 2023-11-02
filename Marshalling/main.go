package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name   string
	Fur string
	Number int
}

func main() {

	c := animal{
		Name:   "Cats",
		Fur: "White",
		Number: 6,
	}
	fmt.Println(c)

	d := animal{
		Name:   "Dogs",
		Fur: "Golden",
		Number: 3,
	}
	fmt.Println(d)

	pet := []animal{c, d}
	fmt.Println(pet)

	bs, err := json.Marshal(pet)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

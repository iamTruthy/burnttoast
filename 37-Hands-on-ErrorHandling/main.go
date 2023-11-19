/* 
instead of using the blank identifier, make sure the code is checking and handling the error.

type person struct {
	First string
	Last string
	Saying []string
}

func main () {

	p := person {
		First: "James",
		Last: "Bond",
		Saying: []string{"I am Bond", "James Bond"},
	}

	bs, _ := json.Marshal(p)
	fmt.Println(string(bs))
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Saying []string
}

func main () {

	p := person {
		First: "James",
		Last: "Bond",
		Saying: []string{"I am Bond", "James Bond"},
	}

	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error: Marshalling unsuccessful",err)
	}
	fmt.Println(string(bs))
	
}
package main

import "fmt"

func main() {

	x := func() {
		m := map[string]int{
			"Jerry": 23,
			"Frank": 25,
		}
		fmt.Println(m)
		fmt.Printf("%v\n", m["Frank"])
	}
	x()

}

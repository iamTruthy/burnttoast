package main

import (
	"fmt"

)

func main() {
	a := 35
	b := 15
	c := 5
	d := 50
	e := 10


	switch {
	case a+b != 50:
		fmt.Println("wrong a+b is equal to 50 try again")
	case a+b == 50:
		fmt.Println("correct! a+b is equal to 50")	
	}

	switch{
	case e == 10 && b-e != c:
		fmt.Println("b-e is equal c")
	case b-e == 5:
		fmt.Println("b-e is equal to 5")
		fallthrough
	case e > c:
		fmt.Println("e is greater than c")
	}

	switch e & c {
	case 10 & 50:
		fmt.Println(!true," e and c is 10 and 5")
	case 50 & 35:
		fmt.Println(!true," e and c is 10 and 5")
	case 10 & 15:
		fmt.Println(!true," e and c is 10 and 5")
	case 10 & 5:
		fmt.Println(!false," e and c is 10 and 5")

	}

	if b > d {
		fmt.Println("wrong ", !true)
	} else if b < d {
		fmt.Println("correct ", b < d)
	}

	if a+b == d && d-b != a {
		fmt.Println("wrong ", !true)
	} else if a+b == d && d-b == a {
		fmt.Println("correct ", !false)
	}

}
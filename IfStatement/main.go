package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := 35
	b := 15
	c := 5
	d := 50

	if a+b != c|d {
		fmt.Println("wrong ", !true)
	} else {
		fmt.Println("correct! answer is ", c, " and", d)
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

	if b*b+c >= a {
		fmt.Println(!false)
	} else if b*b+c != a {
		fmt.Println(!true)
	}

	if b*b+c >= a || b+c == a {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	if e := 3 * rand.Intn(c) ; e >= c {
		fmt.Printf("e is %v and is greater than or equal to %v \n", e,c)
	} else {
		fmt.Printf("e is %v and is less than %v \n", e,c)
	}

}

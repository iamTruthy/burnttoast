package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// a := 15
	z := 9

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)

		switch {
		case x <= 5:
			fmt.Printf("iteration %v \t has a value of %v\n", i, x)
		}
	}

	for {
		fmt.Printf("z is now %v\n", z)
		if z > 30 {
			break
		}
		z++
	}

	/*
	the correct code to print out ODD NUMBERS using the 
	`continue` statement.
	*/

	for b := 0; b < 100; b++ {
		if b%2 != 1 {
			continue
		}
		fmt.Println(b)
	}


	/* 
	another way to print out ODD NUMBERS 
	but without the `continue` statement (used by the instructor)
	*/ 
	for n := 0; n < 100; n++ {
		if n%2 != 0 {
			fmt.Println(n)
		}

	}




	//An infinite Loop
	// for {
	// fmt.Printf("a is now %v\n", a)
	// if a > 50 {
	// }
	// a++
	// }

}

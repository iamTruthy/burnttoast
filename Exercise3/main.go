package main

import (
	"fmt"
	"math/rand"
)

func main() {

	a := 15
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

	// To show Modulus
	for b := 0; b < 100; b++ {
		if b%2 != 1 {
			continue
		}
		fmt.Println(b)
	}

	// An infinite Loop
	for {
		fmt.Printf("a is now %v\n", a)
		if a > 50 {
		}
		a++
	}

}

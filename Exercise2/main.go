package main

import (
	"fmt"
	"math/rand"
)

func init() {

	for j := 0; j < 100; j++ {
		y := rand.Intn(350)

		switch {
		case y <= 100:
			fmt.Println("y is between 0 and 100")
		case y >= 101 && y <= 200:
			fmt.Println("y is between 101 and 200")
		case y >= 201 && y <= 300:
			fmt.Println("y is between 201 and 300")
		case y >= 301 && y <= 350:
			fmt.Println("y is between 301 and 350")
		default:
			fmt.Println("y is more than 350")
		}
	}
}

func main() {

	x := rand.Intn(250)
	a := rand.Intn(10)
	b := rand.Intn(10)

	if x <= 100 {
		fmt.Println("x is between 0 and 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("x is between 101 and 200")
	} else if x >= 201 && x <= 250 {
		fmt.Println("x is between 201 and 250")
	}

	/*
		if a < 4 && b < 4 {
			fmt.Println("a and b are both less than 4")
		} else if a < 6 && b > 6 {
			fmt.Println("a and b are both greater than 6")
		} else if a >= 4 && a <= 6 {
			fmt.Println("a is greater than or equal to 4 and less than or equal to 6")
		} else if b != 5 {
			fmt.Println("b is not equal 5")
		} else {
			fmt.Println("none of the previous cases were met")
		}
	*/

	switch {
	case a < 4 && b < 4:
		fmt.Println("a and b are both less than 4")
	case a < 6 && b > 6:
		fmt.Println("a and b are both greater than 6")
	case a >= 4 && a <= 6:
		fmt.Println("a is greater than or equal to 4 and less than or equal to 6")
	case b != 5:
		fmt.Println("b is not equal 5")
	default:
		fmt.Println("none of the previous cases were met")
	}

	// This will giive the same results as the code above is executed properly //

	for i := 0; i < 100; i++ {
		fmt.Printf("loop %v \n ", i)
	}

}

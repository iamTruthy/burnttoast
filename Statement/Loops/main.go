package main

import "fmt"

func main() {

	a := 15
	b := 10
	c := 9

	for d := 14; d >= c; d-- {
		fmt.Printf("d is greater than or equals %v \n", c)
	}

	for b < a {
		fmt.Printf("%v is less than %v \n", b, a)
		a--
	}

	for {
		if a > 25 {
			fmt.Printf("%v is greater than 25 \n", a)
		} else if a < 25 {
			fmt.Printf("%v is less than 25\n", a)
			break
		}
	}

	for {
		fmt.Printf("%v is greater than or equal 18 \n", c)
		if c > 18 {
			break
		}
		c++
	}

	for x := 4; x <= 35; x++ {
		if x%3 != 0 {
			continue
		}
		fmt.Println(x)
	}

	for i := 4; i < 9; i++ {
		fmt.Println("first loop")
		for x := 3; x < 9; x++ {
			fmt.Println("second loop")
		}
	}

}

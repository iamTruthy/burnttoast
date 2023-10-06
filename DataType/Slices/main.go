package main

import "fmt"

func main () {
	cards := [] string {"Frank","Kmoney","_","Luca"}
	cards = append(cards,"notFibonaccii","MadArmani")
	for i := range cards {
		fmt.Printf("value of i is %v in cards %v \n", i,cards)

	}
}
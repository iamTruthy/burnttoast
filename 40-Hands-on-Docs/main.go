package main

// Package dog is imported
import (
	"fmt"
	"github.com/iamTruthy/dog"
)

func main() {
// Package dog is called in func main and assigned to a variable human. 
// dog.Years takes in a value of type int which is multiplied by whatever value is sent into the function call 

/*
example:

dog.Years(a) =  a * bingo

*/

/*

func Years(a int) int {

	return (a * bingo)
}

*/
	human := dog.Years(6)
	fmt.Printf("human years to dog years is: %v\n",human)

}

package main

// Package dog is imported
import (
	"fmt"
	"github.com/iamTruthy/dog"
)
// const age sets "age" to any value of type int
const age int = 4

func main() {
	// Package dog is called in func main and assigned to a variable human.
	// dog.Years takes in a value of type int which is multiplied by whatever value is sent into the function call

	human := dog.Years(age)
	fmt.Printf("%v human years to dog years is: %v years\n", age, human)

}

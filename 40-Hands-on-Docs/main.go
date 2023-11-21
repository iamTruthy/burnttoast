package main

// Package dog is imported
import (
	"fmt"
	"github.com/iamTruthy/dog"
)

func main() {
// Package dog is called in func main and assigned to a variable human. 
// it takes in a value of type int which is 
/*

func Years(a int) int {

	return (a * bingo)
}

*/
	human := dog.Years(2)
	fmt.Println(human)

}

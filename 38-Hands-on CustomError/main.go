package main

import "fmt"

/*

Create a struct “customErr” which implements the builtin.error interface ( Error() string). Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”.

*/

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("Error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {

	cmEr := customError{
		info: "Transaction declined",
	}
	foo(cmEr)
}

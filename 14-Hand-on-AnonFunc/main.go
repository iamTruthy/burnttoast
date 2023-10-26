package main

import "fmt"

func main() {
	func() {
		fmt.Println(name("Jerry"))
	}()

	func(p string) {
		fmt.Println("I'd love to visit ", p)
	}("Italy")

}

func name(n string) string {
	return fmt.Sprint("My name is ", n)
}

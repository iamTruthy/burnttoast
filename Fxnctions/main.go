package main

import "fmt"

func main() {
	foo()                             // no params, no returns
	mode("Do Not Disturb")            // 1 param, no returns
	s := brand("Le fleur by Lacoste") // 1 param, 1 return
	fmt.Println(s)
	n, t := person("Yoda", "4pm")
	fmt.Println(n, t)

}

func foo() {
	fmt.Println("I am Obiwan Kenobi") // no params, no returns
}

func mode(s string) {
	fmt.Println("My phone was on ", s) // 1 param, no returns
}

func brand(s string) string {
	return fmt.Sprint("Favorite brand is ", s) // 1 param, 1 return
}

func person(name string, time string) (string, string) {
	return fmt.Sprint(name, " has been away since "), time //2 params, 2 returns
}

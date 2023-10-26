package main

import "fmt"

func main() {
	fmt.Println(language("Golang"))
}

func language(lang string) string {
	return fmt.Sprint("My fave programming language is ", lang)
}

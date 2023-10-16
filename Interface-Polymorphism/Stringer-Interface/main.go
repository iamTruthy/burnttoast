package main

import (
	"fmt"
	"strconv"
)

type movie struct {
	title string
}

func (m movie) String() string {
	return fmt.Sprint("The name of the Anime is ", m.title)
}

type rating int

func (r rating) String() string {
	return fmt.Sprint("The rating is ", strconv.Itoa(int(r)))
}

func main() {

	m := movie{
		title: "Blade of The Immortal",
	}

	fmt.Println(m)

	var r rating = 10

	fmt.Println(r)

}

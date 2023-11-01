package main

import (
	"fmt"
	"log"
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

// This is a Wrapper function. 

func logInfo(s fmt.Stringer) {
	log.Println("Log from today: ", s.String())
}

func main() {

	m := movie{
		title: "Blade of The Immortal",
	}

	logInfo(m)

	var r rating = 10

	logInfo(r)

}

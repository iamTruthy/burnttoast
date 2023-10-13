package main

import "fmt"

// This is Using Composit literal to create a struct (compare this code with the "Embedding" folder to se the difference)

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

type engine struct {
	electric bool
}

func main() {

	de := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Chevrolet",
		model: "Impala",
		doors: 2,
		color: "Black",
	}

	fmt.Println(de)
	fmt.Println(de.make)

	ev := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Tesla",
		model: "V",
		doors: 2,
		color: "White",
	}

	fmt.Println(ev)
	fmt.Println(ev.make)

}

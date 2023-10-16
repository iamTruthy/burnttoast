package main

import "fmt"

type artist struct {
	music string
}

func (a artist) song() {
	fmt.Println("i love ", a.music, " music")
}

func main() {
	a := artist{
		music: "Rap",
	}

	a.song()

}

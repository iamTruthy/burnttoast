package main

import "fmt"

type music struct {
	genre  string
	artist string
	song   string
	year   int
}

type award struct {
	music
	grammy bool
}

func main() {

	// This is an example of Embedding a struct
	aw := award{
		music: music{
			genre:  "Afrobreat",
			artist: "Burna Boy",
			song:   "African Giant & Twice As Tall",
			year:   2018,
		},

		grammy: true,
	}

	fmt.Println(aw)
	fmt.Printf("%T\t %#v\n", aw, aw)
	fmt.Println(aw.artist, aw.music.song, aw.year) // This is accessing record by field name with the period notation (tm.year,tm.genre)

	fmt.Println("---------------------------------------------")

	aw2 := award{
		music: music{
			genre:  "Rap",
			artist: "Tylerthecreator",
			song:   "Call Me If You Get Lost: The Estate Sale",
			year:   2023,
		},
		grammy: true,
	}

	fmt.Println(aw2)
	fmt.Printf("%#v \t %T\n", aw2, aw2)
	fmt.Println(aw2.artist, aw2.song, aw2.genre) // This is accessing record by field name with the period notation
	fmt.Println(aw2.music.artist)//This Prints out the same result "aw2.artist" (try it out).


}

package main

import "fmt"

type music struct {
	genre  string
	artist string
	song   string
	year   int
}

func main() {

	tm := music{
		genre:  "Alte",
		artist: "Cruel Santino",
		song:   "Turn Down Miami",
		year:   2018,
	}

	fmt.Println(tm)
	fmt.Printf("%v \t %T\n", tm, tm)
	fmt.Println(tm.artist, tm.song) // This is accessing record by field name with the period notation (tm.year,tm.genre)
	fmt.Println("-----------------------------------")

	tm1 := music{
		genre:  "R&B",
		artist: "Snoh Alegra",
		song:   "I Want You",
		year:   2022,
	}

	fmt.Println(tm1)
	fmt.Printf("%v \t %T\n", tm1, tm1)
	fmt.Println(tm1.artist, tm1.song) // This is accessing record by field name with the period notation
	fmt.Println("-----------------------------------")

	tm2 := music{
		genre:  "Rap",
		artist: "Tylerthecreator",
		song:   "Call Me If You Get Lost: The Estate Sale",
		year:   2023,
	}

	fmt.Println(tm2)
	fmt.Printf("%v \t %T\n", tm2, tm2)
	fmt.Println(tm2.artist, tm2.song) // This is accessing record by field name with the period notation (tm.year,tm.genre)
	fmt.Println("-----------------------------------")

	tm3 := music{
		genre:  "Afrobreat",
		artist: "Burna Boy",
		song:   "African Giant",
		year:   2018,
	}

	fmt.Println(tm3)
	fmt.Printf("%v \t %T\n", tm3, tm3)
	fmt.Println(tm3.artist, tm3.song) // This is accessing record by field name with the period notation (tm.year,tm.genre)
}
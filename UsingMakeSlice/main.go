package main

import "fmt"

func main() {
	as := make([]string, 0)
	as = append(as, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut",
		"Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas",
		"Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota",
		"Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey",
		" New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon",
		" Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas",
		"Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")

	fmt.Println(len(as))
	fmt.Println(cap(as))

	for i := 0; i < len(as); i++ {
		fmt.Println(as[i], i)
		
	}

}

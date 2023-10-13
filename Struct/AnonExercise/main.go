package main

import "fmt"

func main() {

	as := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Tommy",
		friends: map[string]int{
			"Jerry":   23,
			"Maria":   32,
			"Franlin": 27,
		},
		favDrinks: []string{"Monster", "Smirnoff Ice","Water"},
	}

	fmt.Println(as)

	for i, v := range as.friends {
		fmt.Println(i, v)
	}
}

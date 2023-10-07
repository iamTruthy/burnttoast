package main

import "fmt"

func main() {

	cards := []string{"Frank", "Kmoney", "Swoosh", "Luca"}
	cards = append(cards, "Bryan", "MadArmani")
	fmt.Println(cards)

	fmt.Println("-------------")

	for i, v := range cards {
		fmt.Printf("%v \t %v\n", i, v)
	}

	fmt.Println("-------------")

	fmt.Println(len(cards))

	fmt.Println("-------------")

	// accessing items in a slice by index position.
	fmt.Println(cards[0])
	fmt.Println(cards[1])
	fmt.Println(cards[2])
	fmt.Println(cards[3])
	fmt.Println(cards[4])
	fmt.Println(cards[5])

	fmt.Println("------------")

	// this is a more convinient way of accessing items in a slice by index position.
	// prints out same result as the above
	for n := 0; n < len(cards); n++ {
		fmt.Println(cards[n])
	}

	fmt.Println("------------")
	//SLICING A SLICE
	fmt.Println(cards[4:])
	fmt.Println(cards[:2])
	fmt.Println(cards[3:5])
	fmt.Println(cards[0:])
	fmt.Println(cards[:5])

}

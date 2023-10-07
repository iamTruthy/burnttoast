package main

import "fmt"

func main() {

	// Everything in here -- this code block, is an example of an array literal.
	var fe [6]int

	fe[0] = 69
	fmt.Printf("%v\t %T\n", fe, fe)

	x := [...]string{"Goodmorning", "World"}
	fmt.Println(x)

	var i [3]string

	i[0] = "What"
	i[1] = "the Dog"
	i[2] = "do?"

	fmt.Println(i)

	/*
	   Using an array literal to store elements in an array and letting the compiler determine
	   the length of the array, then also print out the array and the length of the array
	*/
	v := [...]string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet",
		"Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter",
		"Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
		"Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar",
		"Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate",
		"Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter",
		"Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry",
		"Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade",
		"Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine",
		"Tracks (GF)"}

	fmt.Println(len(v))

}

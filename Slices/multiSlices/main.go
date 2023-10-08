package main

import "fmt"

func main() {

	//Showing Multidimensional Slices

	pp := []string{"Luca", "Igloo", "Penguins", "Ice"}
	ss := []string{"Wabdoteth", "budzyxbt", "Camol", "Seals"}

	cute := [][]string{pp, ss}

	for i, v := range cute {
		fmt.Println(i, v)

		for x, y := range v {
			fmt.Println(x, y)
		}

	}

}

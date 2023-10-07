package main

import "fmt"

func main() {

	//Showing Multidimensional Slices

	pp := []string{"Luca", "Igloo", "Penguins", "Ice"}
	ss := []string{"Wabdoteth", "budzyxbt", "Camol", "Seals"}
	fmt.Println(pp)
	fmt.Println(ss)

	fmt.Println("----------------------------------")

	cute := [][]string{pp, ss}
	fmt.Println(cute)
}

package main

import "fmt"

func main() {

	m := map[string][]string{

		"Color": {"red", "blue", "biege", "cream", "nude"},
		"Brand": {"Nike", "Lacoste", "Lefeur", "Crocs", "Cookies"},
		"Cars":  {"Chevrolet", "Mustang", "Porsche", "BMW", "Tesla"},
	}

	fmt.Println(m)

	fmt.Println("--------------------------------------")

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Printf("%#v \t %T \t %v\n", k, v, v)

		for i, x := range v {
			fmt.Println(i, x)
		}
	}

	// Deleting a record from the Map

	fmt.Println("-----------------Deleted Map------------------------")

	delete(m, "Color")

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Printf("%#v \t %T \t %v\n", k, v, v)
		for i, x := range v {
			fmt.Println(i, x)
		}
	}


}

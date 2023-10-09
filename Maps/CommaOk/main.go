package main

import "fmt"

func main() {

	m := map[string]string{
		"Jean Piere": "France",
		"Franklin":   "America",
		"Wayne":      "England",
		"Odunsi":     "Nigeria",
	}

	fmt.Println("Jean Piere is from", m["Jean Piere"])
	// This is only printed out because of the Println statement which still access to the map

	fmt.Println("Franklin is from", m["Franklin"])
	fmt.Println("Wayne is from", m["Wayne"])
	fmt.Println("Odunsi is from", m["Odunsi"])

	fmt.Println("------------ for Odunsi ---------------------")

	delete(m, "Jean Piere")

	// Using the comma ok idiom to check if a key "Odunsi" exists
	if w, ok := m["Odunsi"]; !ok {
		fmt.Println(w, "Doesnt exist")
	} else if w, ok := m["Odunsi"]; ok {
		fmt.Println(w, "Exists")
	}

	fmt.Println("---------------------------------")

	for k, v := range m {
		fmt.Println(k, v)

	}

	fmt.Printf("%v\n", m)

	fmt.Println("------------ for Rahul ---------------------")

	nm := make(map[string]int)

	nm["JamesBond"] = 007
	nm["February"] = 14

	// Using the comma ok idom to check if a key "Rahul" exists
	if v, ok := nm["Rahul"]; !ok {
		fmt.Println(v, "Doesn't exist")

	} else{

		fmt.Println("Key Exists")
	}

	fmt.Println("------------------------------")

	fmt.Println("James Bond is Agent", nm["JamesBond"])
	fmt.Println("Someones Birthday is on February", nm["February"])

	fmt.Println("---------------------------------")

	// Deleting a key from a map[]
	delete(nm, "JamesBond")

	for k, v := range nm {
		fmt.Println(k, v)

	}

	fmt.Printf("%v\n", nm)

}

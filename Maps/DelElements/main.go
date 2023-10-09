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


	fmt.Println("---------------------------------")

	delete(m,"Jean Piere")

	 for k, v := range m  {
		fmt.Println(k,v)
		
	}

	fmt.Printf("%v\n", m)

	fmt.Println("---------------------------------")

	nm := make(map[string]int)

		nm["JamesBond"] = 007
		nm["February"] = 14

        fmt.Println("James Bond is Agent", nm["JamesBond"])
		fmt.Println("Someones Birthday is on February", nm["February"])
     
		fmt.Println("---------------------------------")

		// Deleting a key from a map[]
		delete(nm,"JamesBond")

		for k, v := range nm  {
			fmt.Println(k,v)
			
		}
	
		fmt.Printf("%v\n", nm)
	
	

	

}

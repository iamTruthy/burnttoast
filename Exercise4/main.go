package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("i is %v and x is 3\n", i)
		if x := 0 * rand.Intn(5); x == 3 {

		}
	}

	/* create a loop that runs 5 times
	nest a loop within the first loop that also prints 5 times
	print something in each loop to illustrate what is occuring
	*/

	for z := 0; z < 5; z++ {
		fmt.Printf("\n %v first loop running\n", z)
		for k := 0; k < 5; k++ {
			fmt.Printf("\t \t \t %v second loop running\n", k)
		}
	}

	s := []string{"Pugdy Penduins", "MadLads", "Degods", "Azuki", "y00ts"}

	for r, v := range s {
		fmt.Printf("rank %v \t %v\n", r, v)

	}

	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	for v,k := range m {
	fmt.Printf("\nvalue is %v and key is %v\n",k,v)
	}

	rizz := m["James"] 
	fmt.Println(rizz)
	if v, ok := m["James"]; ok {
		fmt.Printf("James found in map with value of %v\n",v)

	}


	yolo :=m["W"]
	fmt.Println(yolo)
	if v, ok := m["yolo"]; !ok {
		fmt.Println("yolo not found in map with value of an int", v)
	}


}

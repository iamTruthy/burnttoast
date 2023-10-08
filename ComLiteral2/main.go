package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for v, t:= range x {
		fmt.Printf("%v - %T - %v\n",t,v,v)
		
	}

	fmt.Println("------------------------------------")

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}

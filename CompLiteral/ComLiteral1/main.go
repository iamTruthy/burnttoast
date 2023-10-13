package main

import "fmt"

func main() {

	var x [5]int

	x[0] = 6
	x[1] = 9
	x[2] = 12
	x[3] = 8
	x[4] = 32

	for v, t:= range x {
		fmt.Printf("%v \t- %T - \t%v\n",t,v,v)
	}

}

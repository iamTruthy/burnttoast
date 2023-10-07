package main

import "fmt"

func main() {

	// DELETING A SLICE

	color := []string{"red", "blue", "brown", "black"}
	fmt.Println(color[:1], color[2:])

	fmt.Println("-------------------")

	card := [...]int{3, 6, 9, 16, 24, 32, 18}
	fmt.Println(card[:2], card[3:])

	fmt.Println("-------------------")

	/*

		Slices are built on top of arrays. A slice is dynamic in that it will grow in size. The underlying
	array, however, does not grow in size. When we create a slice, we can use the built in
	function make to specify how large our slice should be and also how large the underlying
	array should be


	*/

	ns := make([]string, 0, 10)
	fmt.Println(len(ns)) // to show the length of the slice
	fmt.Println(cap(ns)) // to show the cpacity of the slice
	ns = append(ns, "water", "larva", "sunshine", "roses")
	fmt.Println(len(ns))
	fmt.Println(cap(ns))
	fmt.Println(ns)

	fmt.Println("-------------------")
	fmt.Println("-------------------")

	as := make([]int, 0, 10)
	fmt.Println(len(as)) // to show the length of the slice
	fmt.Println(cap(as)) // to show the cpacity of the slice
	as = append(as, 21, 14, 32, 64, 69, 101)
	fmt.Println(len(as))
	fmt.Println(cap(as))
	fmt.Println(as)

	as = append(as, 16, 2, 9, 8, 42)
	fmt.Println(len(as))
	fmt.Println(cap(as))
	fmt.Println(as)

}

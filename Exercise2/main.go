package main

import "fmt"

func main() {

	fmt.Printf("%d \t %b \n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b \n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b \n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b \n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b \n", 1<<5, 1<<5)

	fmt.Printf("%d \t %b \n", 1>>1, 1>>1)
	fmt.Printf("%d \t %b \n", 1>>2, 1>>2)
	fmt.Printf("%d \t %b \n", 1>>3, 1>>3)
	fmt.Printf("%d \t %b \n", 1>>4, 1>>4)
	fmt.Printf("%d \t %b \n", 1>>5, 1>>5)
}

/*

this code shows bit shitfing using Golang. Bit shifting or Bit Operations,
is basically shifting a bit from one position, to another either left or right. This
usually has to do with numeric constants (decimals and binaries).


The code above shifts the bit from left to right and vice versa. this could be achieved using
the greater than (>>) or lessr than sign (<<) depending on where you want to shift the bit to.

*/

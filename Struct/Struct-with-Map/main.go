package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {

	p := person{
		first: "Brandon",
		last:  "Mayhew",
		fav:   []string{"lacoste", "french waltz", "le fleur"},
	}

	fmt.Printf("%v\n", p)

	ap := person{
		first: "Wlater",
		last:  "White",
		fav:   []string{"chocolate", "french vanilla ice cream", "strawberry milkshake"},
	}

	fmt.Printf("%v\n", ap)

	for _, v := range p.fav {
		fmt.Println(p.first, " favorite brands is ", v)
	}

	for _, v := range ap.fav {
		fmt.Println(ap.first, "favorite ice cream flavour is ", v)
	}

	fmt.Println("------------------Struct Map------------------------------------")

	m := map[string]person{

		ap.last: p,
		p.last:  p,
	}

	for _, v := range m {
		fmt.Println(v)
		for _,x := range v.fav {
			fmt.Println(x)
		}
	}

}

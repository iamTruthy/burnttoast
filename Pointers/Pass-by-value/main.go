package main

import "fmt"

func intDelta(id *int) {
	*id = 43
}

func sliceDelta(sd []int) {
	sd[0] = 69

}

func mapDelta(md map[string]int, s string) {
	md[s] = 33

}

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	x := []int{2, 3, 9, 8, 6, 5, 4}
	fmt.Println(x)
	sliceDelta(x)
	fmt.Println(x)
	


	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")
	fmt.Println(m["James"])


}

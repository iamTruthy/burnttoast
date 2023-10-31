package main

import (
	"encoding/json"
	"fmt"
)
// Go to "json to Go" website to convert json to Go when Unmarshalling.
type animal struct {
	Name string `json:"Name"`
	Fur  string `json:"Fur"`
}

func main() {

	s := `[{"Name":"Cats","Fur":"White"},{"Name":"Dogs","Fur":"Golden"}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// pet := []animal{}
	var pet []animal

	err := json.Unmarshal(bs, &pet)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pet)

	for i, v := range pet {
		fmt.Println("\nFavorite Animal", i)
		fmt.Println(v.Name, v.Fur)

	}

}

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`
	bs := []byte(s)

	var users []user

	err := json.Unmarshal(bs, &users)
	if err != nil {
		fmt.Println(err, "Nothing here...")
	}

	fmt.Println(users)
}

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//Will comeback to this later

func main() {

	pswrd := `gibberish1699`

	bs, err := bcrypt.GenerateFromPassword([]byte(pswrd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pswrd)
	fmt.Println(bs)

	loginDetails := `gibberish1699`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginDetails))
	if err != nil {
		fmt.Println(err," Wrong Details, Try Again")
	}
	fmt.Println("Successful")

}
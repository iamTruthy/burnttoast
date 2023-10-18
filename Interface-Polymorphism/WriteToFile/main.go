package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeTo(w io.Writer) {
	w.Write([]byte(p.name))
}

func main() {

	p := person{
		name: "Jenny",
	}

	f, err := os.Create("WriteFile.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	// s := []byte("The writer interface is a fundamental concept used for writing data to various output destinations, including files.")

	// _, err = f.Write(s)
	// if err != nil{
	// log.Fatalf("%s",err)
	// }

	var b bytes.Buffer

	p.writeTo(f)
	p.writeTo(&b)
	fmt.Println(b.String())

}

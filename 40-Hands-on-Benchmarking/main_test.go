package main

import (
	"fmt"
	"github.com/iamTruthy/dog"
	"testing"
)

func TestYears(t *testing.T) {
	human := dog.Years(2)
	fmt.Println(human)
	if human != 14 {
		t.Error("Error")
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}

}

func ExampleYears() {
	fmt.Println(dog.Years(5))
	// Output: 
	// 35
}

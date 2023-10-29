package main

import (
	"fmt"
	"time"
)

func main() {
	timeCount(activity)

}

func activity() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func timeCount(f func()) {
	start := time.Now()
	f()
	stop := time.Since(start)
	fmt.Println(stop)
}

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	variable := 0
	gs := 100

	wg.Add(gs)

	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			v := variable
			v++
			variable = v
			fmt.Println(variable)
			mtx.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("Count:", variable)
}

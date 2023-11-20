package main

import (
	"fmt"
	"log"
)

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (sqE sqrtErr) Error() string {
	return fmt.Sprintf("Math Error: %v %v %v", sqE.lat, sqE.long, sqE.err)
}

func main() {

	err := sqrt(-5)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64 error) {
	if f < 0 {
		// e := errors.New("Wrong Coordinates") --- This could also be use as well. Make sure to import the errors package.
		e := fmt.Errorf("Wrong Coordinates - value was %v", f)
		return sqrtErr{"50.2289 N", "99.4656 W...",e}
	}
	return nil
}

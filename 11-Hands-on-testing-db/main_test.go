package main

import "testing"

func TestAdd(s *testing.T) {
	sum := Add(7, 14)
	if sum != 21 {
		s.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 21)
	}

}

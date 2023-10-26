package main

import "testing"

func TestAdd(s *testing.T) {
	sum := add(42, 16)
	if sum != 58 {
		s.Errorf("Sum was incorrect, got: %d want: %d.", sum, 58)
	}

}

func TestSubtract(d *testing.T) {
	diff := subtract(42, 16)
	if diff != 26 {
		d.Errorf("Difference was incorrect, got: %d want: %d.", diff, 26)
	}
}

func TestDoMath(m *testing.T) {
	math := doMath(42, 16, add)
	if math != 58 {
		m.Errorf("Sum was incorrect, got: %d want: %d.", math, 58)
	}
}

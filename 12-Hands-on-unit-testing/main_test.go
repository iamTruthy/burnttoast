package main

import "testing"

func TestAdd(t *testing.T) {
	sum := add(42, 16)
	if sum != 58 {
		t.Errorf("Sum was incorrect, got: %d want: %d.", sum, 58)
	}

}

func TestSubtract(t *testing.T) {
	diff := subtract(42, 16)
	if diff != 26 {
		t.Errorf("Difference was incorrect, got: %d want: %d.", diff, 26)
	}
}

func TestDoMath(t *testing.T) {
	math := doMath(42, 16, add)
	if math != 58 {
		t.Errorf("Sum was incorrect, got: %d want: %d.", math, 58)
	}
}


// Benchmark measures the performance or speed of func doMath
func BenchmarkDoMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doMath(4, 5, add)
		doMath(4, 5, subtract)

	}
}

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		subtract(4, 5)
	}
}

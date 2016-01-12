package fibonacci

import (
	"testing"
)

var fib = []int{
	0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
	89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946,
}

func testFibonacci(t *testing.T, n int, f func(int) int) {
	expected := fib[n]
	actual := f(n)
	if actual != expected {
		t.Fatalf("Expected %v, but %v:", expected, actual)
	}
}

func TestFibonacciLoop(t *testing.T) {
	// test the function
	for i := range fib {
		testFibonacci(t, i, FibonacciLoop)
	}
}

func TestFibonacciRecursive(t *testing.T) {
	// test the function
	for i := range fib {
		testFibonacci(t, i, FibonacciRecursive)
	}
}

func TestFibonacciRecursiveGoRoutine(t *testing.T) {
	// test the function
	for i := range fib {
		testFibonacci(t, i, FibonacciRecursiveGoRoutine)
	}
}

func TestFibonacciRecursiveGoRoutineLoop(t *testing.T) {
	// test the function
	for i := range fib {
		testFibonacci(t, i, FibonacciRecursiveGoRoutineLoop)
	}
}

func BenchmarkFibonacciLoop(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciLoop(10)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciRecursive(10)
	}
}

func BenchmarkFibonacciRecursiveGoRoutine(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciRecursiveGoRoutine(10)
	}
}

func BenchmarkFibonacciRecursiveGoRoutineLoop(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		FibonacciRecursiveGoRoutineLoop(10)
	}
}

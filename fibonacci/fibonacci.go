package fibonacci

// FibonacciLoop returns a fibonacci value.
// The function is using loop to calculate.
func FibonacciLoop(n int) int {
	if n < 2 {
		return n
	}
	f1, f0 := 1, 0
	fn := f1 + f0
	for i := n; i >= 2; i-- {
		fn = f1 + f0
		f1, f0 = fn, f1
	}
	return fn
}

// FibonacciRecursive returns a fibonacci value
// The function is recursively defined to calculate.
func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

// FibonacciRecursiveGoRoutine returns a fibonacci value
// The function is recursively defined with go routine to calculate.
func FibonacciRecursiveGoRoutine(n int) int {
	if n < 2 {
		return n
	}

	chf1, chf0 := make(chan int), make(chan int)

	go func(f chan int) {
		f <- FibonacciRecursive(n - 1)
	}(chf1)
	go func(f chan int) {
		f <- FibonacciRecursive(n - 2)
	}(chf0)

	f1, f0 := <-chf1, <-chf0

	return f1 + f0
}

// FibonacciRecursiveContinuousGoRoutine returns a fibonacci value
// The function is recursively defined with go routine continuously
// to calculate.
func FibonacciRecursiveContinuousGoRoutine(n int) int {
	if n < 2 {
		return n
	}

	chf1, chf0 := make(chan int), make(chan int)

	go func(f chan int) {
		f <- FibonacciRecursiveContinuousGoRoutine(n - 1)
	}(chf1)
	go func(f chan int) {
		f <- FibonacciRecursiveContinuousGoRoutine(n - 2)
	}(chf0)

	f1, f0 := <-chf1, <-chf0

	return f1 + f0
}

# Fibonacci Function Benchmark

## Check It Out!

```
$ go get -u github.com/eure/go-benchmark/fibonacci
$ cd /path/to/go-benchmark/fibonacci
$ go test -bench . -benchmem
```


## Benchmark

The result of `go test -bench`.

```
$ go test -bench . -benchmem
PASS
BenchmarkFibonacciLoop                          100000000               12.2 ns/op             0 B/op          0 allocs/op
BenchmarkFibonacciRecursive                      3000000               436 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacciRecursiveGoRoutine             1000000              1815 ns/op             192 B/op          2 allocs/op
BenchmarkFibonacciRecursiveGoRoutineLoop           10000            130769 ns/op           16896 B/op        176 allocs/op
ok      github.com/eure/go-benchmark/fibonacci  6.255s
```


## Functions

### FibonacciLoop

Calculate a value with simply for-loop.

```go
func FibonacciLoop(n int) int {
	// ...
	for i := n; i >= 2; i-- {
		// ...
	}
	return
}
```

### FibonacciRecursive

Calculate a value with recursively function

```go
func FibonacciRecursive(n int) int {
	// ...
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
```

### FibonacciRecursiveGoRoutine

Calculate a value with recursively function and go routine

```go
func FibonacciRecursiveGoRoutine(n int) int {
	// ...
	chf1, chf0 := make(chan int), make(chan int)
	go func(f chan int) { f <- FibonacciRecursive(n - 1) }(chf1)
	go func(f chan int) { f <- FibonacciRecursive(n - 2) }(chf0)
	// ...
	return
}
```

### FibonacciRecursiveGoRoutineLoop

Calculate a value with recursively function and continuously go routine

```go
func FibonacciRecursiveGoRoutineLoop(n int) int {
	// ...
	chf1, chf0 := make(chan int), make(chan int)
	go func(f chan int) { f <- FibonacciRecursiveGoRoutineLoop(n - 1) }(chf1)
	go func(f chan int) { f <- FibonacciRecursiveGoRoutineLoop(n - 2) }(chf0)
	// ...
	return
}
```

## Testing

```
$ go test -coverprofile=cover.out . && go tool cover -func=cover.out
ok      github.com/eure/go-benchmark/fibonacci  0.106s  coverage: 100.0% of statements
github.com/eure/go-benchmark/fibonacci/fibonacci.go:5:  FibonacciLoop                   100.0%
github.com/eure/go-benchmark/fibonacci/fibonacci.go:20: FibonacciRecursive              100.0%
github.com/eure/go-benchmark/fibonacci/fibonacci.go:29: FibonacciRecursiveGoRoutine     100.0%
github.com/eure/go-benchmark/fibonacci/fibonacci.go:51: FibonacciRecursiveGoRoutineLoop 100.0%
total:                                                  (statements)                    100.0%
```

## License

- The MIT License (MIT)

## Contributor

- [Shintaro Kaneko](https://github.com/kaneshin)


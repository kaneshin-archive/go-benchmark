# Math Function Benchmark

## Check It Out!

```
$ go get -u github.com/eure/go-benchmark/math
$ cd /path/to/go-benchmark/math
$ go test -bench .
```

## Benchmark

The result of `go test -bench`.

```
$ go test -bench .
PASS
BenchmarkSqrt           30000000                52.5 ns/op
BenchmarkMathSqrt       2000000000               0.35 ns/op
ok      github.com/eure/go-benchmark/math       2.373s
```

## Functions

### Sqrt

Calculate a square root value of x.

```go
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}
```

## Testing

```
$ go test -coverprofile=cover.out . && go tool cover -func=cover.out
ok      github.com/eure/go-benchmark/math       0.002s  coverage: 100.0% of statements
github.com/eure/go-benchmark/math/sqrt.go:3:    Sqrt            100.0%
total:                                          (statements)    100.0%
```

## License

- The MIT License (MIT)

## Contributor

- [Shintaro Kaneko](https://github.com/kaneshin)


package math

import (
	"math"
	"testing"
)

func BenchmarkSqrt(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		Sqrt(2)
	}
}

func BenchmarkMathSqrt(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		math.Sqrt(2)
	}
}

// TestSqrt ...
func TestSqrt(t *testing.T) {

	// eps is for measurement
	const eps = 1.0e-5

	type candidate struct {
		input    float64
		expected float64
	}
	candidates := []candidate{
		{1.0, 1.0},
		{2.0, math.Sqrt2},
		{
			math.E,     // = 2.71828182845904523536028747135266249775724709369995957496696763
			math.SqrtE, // = 1.64872127070012814684865078781416357165377610071014801157507931
		},
		{
			math.Pi,     // = 3.14159265358979323846264338327950288419716939937510582097494459
			math.SqrtPi, // = 1.77245385090551602729816748334114518279754945612238712821380779
		},
		{
			math.Phi,     // = 1.61803398874989484820458683436563811772030917980576286213544862
			math.SqrtPhi, // = 1.27201964951406896425242246173749149171560804184009624861664038
		},
	}

	for _, c := range candidates {
		approx := Sqrt(c.input)
		if math.Abs(approx-c.expected) > eps {
			t.Fatalf("Expected %v, but %v:", c.expected, approx)
		}
	}
}

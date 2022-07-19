package math

import (
	"math"
	"testing"
)

var e = math.E
var pi = math.Pi
var result float64

func BenchmarkAbs(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Abs(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkAcos(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Acos(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkAsin(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Asin(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkAtan(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Atan(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkAtan2(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Atan2(e, pi)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkCbrt(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Cbrt(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkCeil(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Ceil(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkCos(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Cos(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkCosh(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Cosh(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkExp(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Exp(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkFloor(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Floor(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkLog(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Log(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkMax(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Max(e, pi)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkMin(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Min(e, pi)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkPow(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Pow(e, pi)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkRound(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Round(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkSin(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Sin(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkSinh(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Sinh(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkSqrt(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Sqrt(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkTan(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Tan(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkTanh(b *testing.B) {
	var r float64
	for n := 0; n < b.N; n++ {
		r = math.Tanh(e)
	}

	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

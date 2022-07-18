package baseline

import (
	"testing"
)

func NoOp() {}

func BenchmarkNothing(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NoOp()
	}
}

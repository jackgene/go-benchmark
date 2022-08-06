package copy

import (
	"testing"
)

type Data struct {
	Bytes [128 * 1024 * 1024]int8
}

func BenchmarkCopy(b *testing.B) {
	var data1 Data
	var data2 Data
	for n := 0; n < b.N; n++ {
		if n%2 == 0 {
			data1 = data2
		} else {
			data2 = data1
		}
	}
}

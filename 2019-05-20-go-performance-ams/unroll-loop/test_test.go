package main

import (
	"testing"
)

var (
	fls = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum float64
)

func init() {
	// 2^5 * 9 = 288 floats
	for i := 0; i < 5; i++ {
		fls = append(fls, fls...)
	}
	// fls = append(fls, 1, 2, 3) // adding 1,2 or 3 floats isn't really going to affect performance.
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum += sumFloat64(fls)
	}
}

func BenchmarkSumUnrolled(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum += sumFloat64Unrolled(fls)
	}
}

func BenchmarkSumUnrolled2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sum += sumFloat64Unrolled2(fls)
	}
}

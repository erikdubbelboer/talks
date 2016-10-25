package test

import (
	"testing"
)

type big struct {
	a, b, c, d, e, f, g, h float64
}

var s *big
var h big

func forceHeap(x *big) {
	s = x
}

func forceStack(x big) {
	h = x
}

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big{}

		forceHeap(&x)
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := big{}

		forceStack(x)
	}
}

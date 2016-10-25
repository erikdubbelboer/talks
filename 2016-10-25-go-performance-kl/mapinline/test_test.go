package test

import (
	"testing"
)

type big struct {
	// Go inlines keys and values smaller than 128 bytes.
	// See: https://github.com/golang/go/blob/5d8324e6822e34c900e36d67adf640cee6693d25/src/runtime/hashmap.go#L70-L75
	a [20]float64
}

var g1 map[int]int
var g2 map[int]big

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := make(map[int]int)
		x[1] = 2
		x[3] = 4
		g1 = x
	}
}

func BenchmarkBig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := make(map[int]big)
		x[1] = big{}
		x[3] = big{}
		g2 = x
	}
}

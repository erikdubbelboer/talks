package test

import (
	"testing"
)

type big struct {
	a, b, s, d, e, f, g, h float64
}

func byValue(x big) {
	_ = x.h
}

func byRef(x *big) {
	_ = x.h
}

func BenchmarkByValue(b *testing.B) {
	x := big{}
	for i := 0; i < b.N; i++ {
		byValue(x)
	}
}

func BenchmarkByRef(b *testing.B) {
	x := big{}
	for i := 0; i < b.N; i++ {
		byRef(&x)
	}
}

package main

import (
	"testing"
	"unsafe"
)

type X struct {
	a [14]int64 // 112

	b int32 // 4
	c float64 // 8
	d int32 // 4
}

type Y struct {
	a[14]int64

	b int32
	d int32
	c float64
}

func TestXIsNotTooBig(t *testing.T) {
	if s := unsafe.Sizeof(X{}); s > 128 {
		t.Fatalf("X is %v bytes", s)
	}
}

func BenchmarkWrong(b *testing.B) {
	m := make(map[int]X)

	for i := 0; i < b.N; i++ {
		m[i] = X{}
	}
}

func BenchmarkRight(b *testing.B) {
	m := make(map[int]Y)

	for i := 0; i < b.N; i++ {
		m[i] = Y{}
	}
}

func BenchmarkRighter(b *testing.B) {
	m := make(map[int]Y, b.N)

	for i := 0; i < b.N; i++ {
		m[i] = Y{}
	}
}

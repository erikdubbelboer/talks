package main

import (
	"testing"
	"unsafe"
)

type TooBig struct {
	a [14]int64 // 112 (14*8) bytes

	b int32   // 4
	c float64 // 8
	d int32   // 4
}

type JustRight struct {
	a [14]int64 // 112 (14*8) bytes

	b int32   // 4
	d int32   // 4
	c float64 // 8
}

func TestTooBigIsNotTooBig(t *testing.T) {
	if s := unsafe.Sizeof(TooBig{}); s > 128 {
		t.Fatalf("TooBig is %v bytes", s)
	}
}

func BenchmarkTooBig(b *testing.B) {
	m := make(map[int]TooBig)

	for i := 0; i < b.N; i++ {
		m[i] = TooBig{}
	}
}

func BenchmarkJustRight(b *testing.B) {
	m := make(map[int]JustRight)

	for i := 0; i < b.N; i++ {
		m[i] = JustRight{}
	}
}

func BenchmarkJustRightPrealloc(b *testing.B) {
	m := make(map[int]JustRight, b.N) // HL

	for i := 0; i < b.N; i++ {
		m[i] = JustRight{}
	}
}

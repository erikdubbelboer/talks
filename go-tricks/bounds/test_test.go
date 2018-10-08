package main

import (
	"testing"
)

//go:noinline
func Without(b []byte, v uint64) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

//go:noinline
func With(b []byte, v uint64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
}

//go:noinline
func Different(b []byte, v uint64) {
	b[7] = byte(v >> 56)
	b[6] = byte(v >> 48)
	b[5] = byte(v >> 40)
	b[4] = byte(v >> 32)
	b[3] = byte(v >> 24)
	b[2] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[0] = byte(v)
}

func BenchmarkWithout(b *testing.B) {
	buf := make([]byte, 20)

	for i := 0; i < b.N; i++ {
		Without(buf, uint64(i))
	}
}

func BenchmarkWith(b *testing.B) {
	buf := make([]byte, 20)

	for i := 0; i < b.N; i++ {
		With(buf, uint64(i))
	}
}

func BenchmarkDifferent(b *testing.B) {
	buf := make([]byte, 20)

	for i := 0; i < b.N; i++ {
		Different(buf, uint64(i))
	}
}

// https://golang.org/test/checkbce.go

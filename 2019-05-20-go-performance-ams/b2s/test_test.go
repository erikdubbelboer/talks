package main

import (
	"strconv"
	"testing"
	"unsafe"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkSlow(b *testing.B) {
	input := []byte{0x31, 0x33, 0x33, 0x37}

	for i := 0; i < b.N; i++ {
		strconv.Atoi(string(input))
	}
}
func BenchmarkFast(b *testing.B) {
	input := []byte{0x31, 0x33, 0x33, 0x37}

	for i := 0; i < b.N; i++ {
		strconv.Atoi(b2s(input))
	}
}

package main

import (
	"strconv"
	"testing"
)

var (
	base = []byte("foo = ")

	Sink []byte
)

func Itoa() {
	s := make([]byte, 0, 32)
	s = append(s, base...)
	s = append(s, strconv.Itoa(1337)...)
	Sink = s
}

func AppendInt() {
	s := make([]byte, 0, 32)
	s = append(s, base...)
	s = strconv.AppendInt(s, 1337, 10) // HL
	Sink = s
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Itoa()
	}
}

func BenchmarkAppendInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AppendInt()
	}
}

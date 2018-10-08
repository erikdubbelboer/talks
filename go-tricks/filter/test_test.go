package main

import (
	"bytes"
	"testing"
)

func filter1(b []byte, n byte) []byte {
	x := make([]byte, 0)
	for _, v := range b {
		if v != n {
			x = append(x, v)
		}
	}
	return x
}

func filter2(b []byte, n byte) []byte {
	x := make([]byte, 0, len(b))
	for _, v := range b {
		if v != n {
			x = append(x, v)
		}
	}
	return x
}

func filter3(b []byte, n byte) []byte {
	x := b[:0]
	for _, v := range b {
		if v != n {
			x = append(x, v)
		}
	}
	return x
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		r := filter1(a, 3)
		if !bytes.Equal(r, []byte{1, 2, 4, 5, 6, 7, 8, 9, 10}) {
			panic("wrong")
		}
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		r := filter2(a, 3)
		if !bytes.Equal(r, []byte{1, 2, 4, 5, 6, 7, 8, 9, 10}) {
			panic("wrong")
		}
	}
}

func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		r := filter3(a, 3)
		if !bytes.Equal(r, []byte{1, 2, 4, 5, 6, 7, 8, 9, 10}) {
			panic("wrong")
		}
	}
}

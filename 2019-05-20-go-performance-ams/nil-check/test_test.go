package main

import "testing"

type X struct {
	i [64]int
}

var (
	Index = 32
)

func (a *X) doSlow() int {
	r := 0
	for i := 0; i < 10000; i++ {
		if Index+i > 0 {
			r += a.i[Index]
		}
	}
	return r
}

func (a *X) doFast() int {
	_ = a.i

	r := 0
	for i := 0; i < 10000; i++ {
		if Index+i > 0 {
			r += a.i[Index]
		}
	}
	return r
}

func BenchmarkSlow(b *testing.B) {
	a := &X{}
	for i := 0; i < b.N; i++ {
		a.doSlow()
	}
}

func BenchmarkFast(b *testing.B) {
	a := &X{}
	for i := 0; i < b.N; i++ {
		a.doFast()
	}
}

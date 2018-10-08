package main

import (
	"testing"
	"unsafe"
)

//go:nosplit
func noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}

type X struct {
	x *X
	i int
}

//go:noinline
func (a *X) foo(i int) {
	if a.x == nil {
		a.x = a
	}

	a.i = i
}

//go:noinline
func (a *X) bar(i int) {
	if a.x == nil {
		a.x = (*X)(noescape(unsafe.Pointer(a)))
	}

	a.i = i
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x X
		x.foo(i)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x X
		x.bar(i)
	}
}

// https://github.com/golang/go/blob/ce58a39fca067a19c505220c0c907ccf32793427/src/strings/builder.go#L32-L43

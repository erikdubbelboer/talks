package main

import (
	"testing"
)

var Sink int
var True = false

//go:noinline
func cleanup(x *int) {
	xx := 2
	if x != nil {
		xx = *x
	}

	Sink *= xx
}

//go:noinline
func work() {
	Sink *= 2
}

//go:noinline
func noDefer() {
	work()
	cleanup(nil)
}

//go:noinline
func doDefer() {
	defer cleanup(nil)
	work()
}

//go:noinline
func slowDefer() {
	d := 0
	defer func() {
		if True {
			cleanup(&d)
		}
	}()

	work()

	if Sink > 0 {
		d = 10
	}
}

func BenchmarkNoDefer(b *testing.B) {
	True = false
	for i := 0; i < b.N; i++ {
		noDefer()
	}
}

func BenchmarkDoDefer(b *testing.B) {
	True = true
	for i := 0; i < b.N; i++ {
		doDefer()
	}
}

func BenchmarkSlowDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slowDefer()
	}
}

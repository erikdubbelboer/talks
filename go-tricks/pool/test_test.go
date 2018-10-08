package main

import (
	"sync"
	"testing"
)

var (
	badPool = sync.Pool{
		//go:noinline
		New: func() interface{} {
			return make([]byte, 0, 32)
		},
	}

	goodPool = sync.Pool{
		//go:noinline
		New: func() interface{} {
			x := make([]byte, 0, 32)
			return &x
		},
	}
)

func BenchmarkBad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := badPool.Get().([]byte)

		x = append(x, 1)

		badPool.Put(x)
	}
}

func BenchmarkGood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		y := goodPool.Get().(*[]byte)
		x := *y

		x = append(x, 1)

		*y = x
		goodPool.Put(y)
	}
}

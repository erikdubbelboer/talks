package test

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := bytes.Buffer{}

		for j := 0; j < 100; j++ {
			b.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}
	}
}

var (
	p = sync.Pool{
		New: func() interface{} {
			return bytes.Buffer{}
		},
	}
)

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := p.Get().(bytes.Buffer)
		b.Reset()

		for j := 0; j < 100; j++ {
			b.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		}

		p.Put(b)
	}
}

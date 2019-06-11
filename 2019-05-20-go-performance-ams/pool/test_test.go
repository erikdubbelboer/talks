package test

import (
	"bytes"
	"sync"
	"testing"
)

func BenchmarkNoPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := &bytes.Buffer{}
			for j := 0; j < 1000; j++ {
				bb.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			}
		}
	})
}

func BenchmarkPool(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := p.Get().(*bytes.Buffer)
			bb.Reset() // Don't forget to reset the contents!
			for j := 0; j < 1000; j++ {
				bb.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			}
			p.Put(bb)
		}
	})
}

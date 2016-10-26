package test

import (
	"bytes"
	"sync"
	"testing"
)

const work = 10000

// No pool

func BenchmarkNoPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := &bytes.Buffer{}
			for j := 0; j < work; j++ {
				bb.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			}
		}
	})
}

// sync.Pool

func BenchmarkSyncPool(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := p.Get().(*bytes.Buffer)
			bb.Reset()
			for j := 0; j < work; j++ {
				bb.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			}
			p.Put(bb)
		}
	})
}

// Channel pool

type Pool struct {
	New func() interface{}

	pool chan interface{}
}

func NewPool(n func() interface{}, size int) Pool {
	p := Pool{
		New:  n,
		pool: make(chan interface{}, size),
	}

	return p
}

func (p Pool) Get() interface{} {
	select {
	case o := <-p.pool:
		return o
	default:
		return p.New()
	}
}

func (p Pool) Put(o interface{}) {
	select {
	case p.pool <- o:
	default:
	}
}

func BenchmarkChanPool(b *testing.B) {
	p := NewPool(func() interface{} {
		return &bytes.Buffer{}
	}, 32)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bb := p.Get().(*bytes.Buffer)
			bb.Reset()
			for j := 0; j < work; j++ {
				bb.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
			}
			p.Put(bb)
		}
	})
}

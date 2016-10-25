package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"unsafe"
)

func BenchmarkPointer(b *testing.B) {
	var ptr unsafe.Pointer

	m := make(map[int]int)
	m[1] = 2

	atomic.StorePointer(&ptr, unsafe.Pointer(&m))

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := (*map[int]int)(atomic.LoadPointer(&ptr))
			_ = (*m)[1]
		}
	})
}

func BenchmarkValue(b *testing.B) {
	var ptr atomic.Value

	m := make(map[int]int)
	m[1] = 2

	ptr.Store(m)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			m := ptr.Load().(map[int]int)
			_ = m[1]
		}
	})
}

func BenchmarkRWMutex(b *testing.B) {
	var l sync.RWMutex

	m := make(map[int]int)
	m[1] = 2

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.RLock()
			_ = m[1]
			l.RUnlock()
		}
	})
}

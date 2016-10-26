package test

// go test -bench . -benchmem -benchtime 5s
// Show CPU usage.

import (
	"sync"
	"sync/atomic"
	"testing"
)

func doit() {
	for i := 0; i < 1000000; i++ {
	}
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doit()
	}
}

func BenchmarkWaitGroup(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			doit()
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkChannel(b *testing.B) {
	r := make(chan struct{})
	for i := 0; i < b.N; i++ {
		go func() {
			doit()
			r <- struct{}{}
		}()
	}
	for i := 0; i < b.N; i++ {
		<-r
	}
}

func BenchmarkAtomic(b *testing.B) {
	waiting := int64(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			doit()
			atomic.AddInt64(&waiting, -1)
		}()
	}
	for atomic.LoadInt64(&waiting) > 0 {
	}
}

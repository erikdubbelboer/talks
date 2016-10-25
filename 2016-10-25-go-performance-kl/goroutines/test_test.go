package test

// go test -bench . -benchmem -benchtime 5s
// Show CPU usage.

import (
	"sync"
	"testing"
)

func doit() {
	for i := 0; i < 100000; i++ {
	}
}

func BenchmarkSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doit()
	}
}

func BenchmarkParallel(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			doit()
		}()
	}
	wg.Wait()
}

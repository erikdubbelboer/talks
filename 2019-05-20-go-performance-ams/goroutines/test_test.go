package test

import (
	"sync"
	"testing"
	"time"
)

func work() int {
	r := 0
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond)
		r += i
	}
	return r
}

func BenchmarkSerial(b *testing.B) {
	r := 0
	for i := 0; i < b.N; i++ {
		r += work()
	}
	_ = r
}

func BenchmarkWaitGroup(b *testing.B) {
	r := 0
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(b.N)

	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()

			t := work()
			mu.Lock()
			r += t
			mu.Unlock()
		}()
	}
	wg.Wait()

	_ = r
}

func BenchmarkChannel(b *testing.B) {
	r := 0
	c := make(chan int)
	for i := 0; i < b.N; i++ {
		go func() {
			c <- work()
		}()
	}
	for i := 0; i < b.N; i++ {
		r += <-c
	}
	_ = r
}

package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

//go:noinline
func count() int64 {
	var r int64
	for i := int64(0); i < 1000000; i++ {
		r += i % 10
	}
	return r
}

func countParallel() int64 {
	var r int64
	var wg sync.WaitGroup
	wg.Add(10)
	for n := 0; n < 10; n++ {
		go func() {
			for i := int64(0); i < 100000; i++ {
				atomic.AddInt64(&r, i%10)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return r
}

func countSharded() int64 {
	var r [10]int64
	var wg sync.WaitGroup
	wg.Add(10)
	for n := 0; n < 10; n++ {
		go func(n int) {
			for i := int64(0); i < 100000; i++ {
				r[n] += i % 10
			}
			wg.Done()
		}(n)
	}
	wg.Wait()
	return r[0] + r[1] + r[2] + r[3] + r[4] + r[5] + r[6] + r[7] + r[8] + r[9]
}

func countShardedPadding() int64 {
	var r [10]struct {
		r int64
		_ [7]int64
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for n := 0; n < 10; n++ {
		go func(n int) {
			for i := int64(0); i < 100000; i++ {
				r[n].r += i % 10
			}
			wg.Done()
		}(n)
	}
	wg.Wait()
	return r[0].r + r[1].r + r[2].r + r[3].r + r[4].r + r[5].r + r[6].r + r[7].r + r[8].r + r[9].r
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count()
	}
}

func BenchmarkCountParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countParallel()
	}
}

func BenchmarkCountSharded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSharded()
	}
}

func BenchmarkCountShardedPadding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countShardedPadding()
	}
}

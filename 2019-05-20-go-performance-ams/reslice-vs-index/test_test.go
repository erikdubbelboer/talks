package main

import (
	"runtime"
	"testing"
)

type summer struct {
	s []int
	i int
}

func (s *summer) reslice() int {
	if len(s.s) == 0 {
		return 0
	}

	r := s.s[0]
	s.s = s.s[1:]

	return r + s.reslice()
}

func (s *summer) index() int {
	if len(s.s) == s.i {
		return 0
	}

	r := s.s[s.i]
	s.i++

	return r + s.index()
}

var (
	is  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum int
)

func init() {
	// 2^5 * 9 = 288 ints
	for i := 0; i < 5; i++ {
		is = append(is, is...)
	}

	_ = runtime.GC
	go func() {
		for {
			runtime.GC()
		}
	}()
}

func BenchmarkReslice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := summer{s: is}
		sum += s.reslice()
	}
}

func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := summer{s: is}
		sum += s.index()
	}
}

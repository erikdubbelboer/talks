package main // OMIT

import "testing" // OMIT

func appendItems(s []int) []int {
	for i := 0; i < 10000; i++ {
		s = append(s, i)
	}
	return s
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendItems(make([]int, 0))
	}
}

func BenchmarkPreallocate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendItems(make([]int, 0, 10000)) // HL
	}
}

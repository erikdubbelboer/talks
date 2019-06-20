package test // OMIT

import "testing" // OMIT

var heapSink *int64
var stackSink int64

func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			x := int64(i)
			heapSink = &x
		}
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			x := int64(i)
			stackSink = x
		}
	}
}

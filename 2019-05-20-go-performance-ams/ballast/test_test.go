package main

import "testing"

type big struct {
	a, b, c, d *int
}

var Sink *big
var Ballast []byte

func somethingGeneratingGarbage() {
	for j := 0; j < 1000; j++ {
		Sink = &big{}
	}
}

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		somethingGeneratingGarbage()
	}
}

func BenchmarkBallast(b *testing.B) {
	Ballast = make([]byte, 10*1024*1024)

	for i := 0; i < b.N; i++ {
		somethingGeneratingGarbage()
	}
}

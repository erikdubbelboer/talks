package main

import (
	"unsafe"
)

func sumFloat64(s []float64) float64 {
	a := float64(0)
	for _, f := range s {
		a += f
	}
	return a
}

func sumFloat64Unrolled(s []float64) float64 {
	var a1, a2, a3, a4 float64
	l := len(s)
	lm := l % 4

	for i := lm; i < l; i += 4 {
		a1 += s[i] // HL
		a2 += s[i+1] // HL
		a3 += s[i+2] // HL
		a4 += s[i+3] // HL
	}
	if lm == 1 {
		a1 += s[0]
	} else if lm == 2 {
		a1 += s[0]
		a2 += s[1]
	} else if lm == 3 {
		a1 += s[0]
		a2 += s[1]
		a3 += s[2]
	}

	return a1 + a2 + a3 + a4
}

func sumFloat64Unrolled2(s []float64) float64 {
	var a1, a2, a3, a4 float64

	l := len(s)
	lm := l % 4

	for i := lm; i < l; i += 4 {
		ss := (*[4]float64)(unsafe.Pointer(&s[i]))
		a1 += ss[0]
		a2 += ss[1]
		a3 += ss[2]
		a4 += ss[3]
	}

	if lm == 1 {
		a1 += s[0]
	} else if lm == 2 {
		a1 += s[0]
		a2 += s[1]
	} else if lm == 3 {
		a1 += s[0]
		a2 += s[1]
		a3 += s[2]
	}

	return a1 + a2 + a3 + a4
}

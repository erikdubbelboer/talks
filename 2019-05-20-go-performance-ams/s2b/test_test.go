package main

import (
	"io/ioutil"
	"reflect"
	"testing"
	"unsafe"
)

func s2b(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BenchmarkSlow(b *testing.B) {
	input := "this is a test"

	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write([]byte(input))
	}
}

func BenchmarkFast(b *testing.B) {
	input := "this is a test"

	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write(s2b(input))
	}
}

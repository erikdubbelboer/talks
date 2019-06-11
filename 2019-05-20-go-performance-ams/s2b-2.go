package main

import (
	"os"
	"reflect"
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

func main() {
	s := "this is a test"
	os.Stdout.Write(s2b(s))
}

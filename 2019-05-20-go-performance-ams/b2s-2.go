package main

import (
	"strconv"
	"unsafe"
)

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	have := []byte{0x31, 0x33, 0x33, 0x37}

	n, _ := strconv.Atoi(b2s(have))
}

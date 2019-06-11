// +build ignore OMIT

package main

import (
	"os"
	"strconv"
)

var (
	base = []byte("foo = ")
)

func main() {
	s := make([]byte, 0, 32)
	s = append(s, base...)
	s = strconv.AppendInt(s, 1337, 10) // HL
	os.Stdout.Write(s)
}

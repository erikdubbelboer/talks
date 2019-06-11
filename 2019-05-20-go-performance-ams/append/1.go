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
	s = append(s, strconv.Itoa(1337)...)
	os.Stdout.Write(s)
}

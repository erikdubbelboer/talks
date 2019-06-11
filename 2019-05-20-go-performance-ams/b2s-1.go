package main

import (
	"strconv"
)

func main() {
	have := []byte{0x31, 0x33, 0x33, 0x37}

	n, _ := strconv.Atoi(string(have))
}

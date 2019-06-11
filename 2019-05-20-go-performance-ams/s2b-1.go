package main

import (
	"os"
)

func main() {
	s := "this is a test"
	os.Stdout.Write([]byte(s))
}

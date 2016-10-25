package main

import (
	"fmt"
	"time"
)

var g int

func show() {
	for {
		fmt.Println(g)

		time.Sleep(time.Second)
	}
}

func change() {
	for {
		g++

		time.Sleep(time.Millisecond)
	}
}

func main() {
	go show()
	go change()

	select {}
}

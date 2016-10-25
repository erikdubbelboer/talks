package main

import (
	"fmt"
	"math/rand"
	"time"
)

var g = make(map[int]int)

func show() {
	for {
		fmt.Println(g[1])

		time.Sleep(time.Second / 10)
	}
}

func change() {
	for {
		if rand.Intn(100) < 50 {
			delete(g, 1)
		} else {
			g[1]++
		}
	}
}

func main() {
	go show()
	go change()

	select {}
}

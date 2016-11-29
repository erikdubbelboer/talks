package main

import (
	"time"
)

func block() {
	for {
	}
}

func p() {
	for {
		println("here")
		time.Sleep(time.Second)
	}
}

var memSink interface{}

func alloc() {
	for {
		time.Sleep(time.Second)

		for i := 0; i < 1024; i++ {
			memSink = make([]byte, 1024)
		}
	}
}

func main() {
	go p()
	go alloc()

	// Starting this goroutine will prevent the scheduler from
	// preempting all go-routines and running the stop-the-world
	// garbage collector cycle. Preventing all other go-routines
	// from waking up and waiting for the stop-the-world forever.
	// See: https://github.com/golang/go/issues/10958
	go block()

	select {}
}

package main

import (
	"sync/atomic"
	"time"
)

func wrong() {
	count := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				time.Sleep(time.Millisecond * 100)
				count++
			}
		}()
	}

	for {
		time.Sleep(time.Second)
		println(count)
	}
}

func good() {
	count := int64(0)

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				time.Sleep(time.Millisecond * 100)
				atomic.AddInt64(&count, 1) // HL
			}
		}()
	}

	for {
		time.Sleep(time.Second)
		println(atomic.LoadInt64(&count)) // HL
	}
}

func main() {
	wrong()
	//good()
}

package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	_ "github.com/erikdubbelboer/bpprof"
	_ "net/http/pprof"
)

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
	runtime.MemProfileRate = 1

	go alloc()
	log.Printf("open http://0.0.0.0:6060/debug/bpprof/heap?debug=1")
	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}

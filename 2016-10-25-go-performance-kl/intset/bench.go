// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const minSize = 20
const maxSize = 50
const tests = 1000000
const findMult = 100 // 100 = 1% hit rate, 1 = 100% hit rate.

type lst []int

func (l lst) has(xxxx int) bool {
	for _, i := range l {
		if i == xxxx {
			return true
		}
	}
	return false
}

type mp map[int]struct{}

func (m mp) has(xxxx int) bool {
	_, ok := m[xxxx]
	return ok
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("size\tlist\tmap\tdiff\n")

	for size := minSize; size <= maxSize; size++ {
		l := make(lst, 0, size)
		m := make(mp, size)
		for i := 0; i < size; i++ {
			e := rand.Int()
			l = append(l, e)
			m[e] = struct{}{}
		}

		found := 0
		start := time.Now()
		for x := 0; x < tests; x++ {
			xxxx := rand.Intn(size * findMult)
			if l.has(xxxx) {
				found++
			}
		}
		listDuration := time.Now().Sub(start)

		start = time.Now()
		for x := 0; x < tests; x++ {
			xxxx := rand.Intn(size * findMult)
			if m.has(xxxx) {
				found++
			}
		}
		mapDuration := time.Now().Sub(start)

		fmt.Printf("%d\t%dms\t%dms\t%dms\n", size, listDuration/time.Millisecond, mapDuration/time.Millisecond, (listDuration-mapDuration)/time.Millisecond)
	}
}

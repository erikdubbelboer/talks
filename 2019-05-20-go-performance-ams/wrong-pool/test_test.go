package test

import (
	"sync"
	"testing"
)

type A struct {
	i int64
}

func TestDoesNotAllocate(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return &A{}
		},
	}

	if allocs := testing.AllocsPerRun(100, func() {
		a := p.Get().(*A)
		a.i = 100
		p.Put(a)
	}); allocs != 0 {
		t.Fatalf("%v allocations!", allocs)
	}
}


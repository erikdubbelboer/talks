package test

import (
	"testing"
)

type Something struct {
	X string
}

func asStruct(a Something) {
	if a.X == "" {
		_ = a.X
	} else {
		panic("...")
	}
}

func asInterface(a interface{}) {
	if s, ok := a.(Something); ok {
		_ = s.X
	} else {
		panic("not A")
	}
}

func BenchmarkStruct(b *testing.B) {
	a := Something{}
	for i := 0; i < b.N; i++ {
		asStruct(a)
	}
}

func BenchmarkIntrface(b *testing.B) {
	a := Something{}
	for i := 0; i < b.N; i++ {
		asInterface(a)
	}
}

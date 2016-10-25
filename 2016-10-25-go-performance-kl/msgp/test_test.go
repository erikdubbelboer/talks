package test

import (
	"encoding/json"
	"testing"
)

type Something struct {
	A float64
	B int64
	C string
}

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Something{}
		if _, err := json.Marshal(x); err != nil {
			panic(err)
		}
	}
}

func BenchmarkFFJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Something{}
		if _, err := x.MarshalJSON(); err != nil {
			panic(err)
		}
	}
}

func BenchmarkMSGP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Something{}
		if _, err := x.MarshalMsg(nil); err != nil {
			panic(err)
		}
	}
}

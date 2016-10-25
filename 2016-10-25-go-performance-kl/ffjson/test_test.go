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

var str = []byte(`{"A":1.2,"B":3,"C":"4"}`)

func BenchmarkJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Something{}
		if err := json.Unmarshal(str, &x); err != nil {
			panic(err)
		}
	}
}

func BenchmarkFFJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := Something{}
		if err := x.UnmarshalJSON(str); err != nil {
			panic(err)
		}
	}
}

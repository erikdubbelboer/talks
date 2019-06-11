type X struct {
	a [14]int64 // 112

	b int32 // 4
	c float64 // 8
	d int32 // 4
}

type Y struct {
	a[14]int64

	b int32
	d int32
	c float64
}

func TestXIsNotTooBig(t *testing.T) {
	if s := unsafe.Sizeof(X{}); s > 128 {
		t.Fatalf("X is %v bytes", s)
	}
}

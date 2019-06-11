type X struct {
	i [64]int
}

var Index int

func (a *X) foo() int {
	_ = a.i // HL

	r := 0
	for i := 0; i < 10000; i++ {
		if Index+i > 0 {
			r += a.i[Index]
		}
	}
	return r
}

pool := sync.Pool{
	New: func() interface{} {
		b := make([]byte, 0)
		return &b
	},
}

sp := pool.Get().(*[]byte)
s := *sp

// ...

// This line is very important.
// Doing `sp = &s` or `pool.Put(&s)` here instead would still cause an allocation.
*sp = s // HL
pool.Put(sp)

func BenchmarkX(b *testing.B) {
	m := make(map[int]X)

	for i := 0; i < b.N; i++ {
		m[i] = X{}
	}
}

func BenchmarkY(b *testing.B) {
	m := make(map[int]Y, b.N) // HL

	for i := 0; i < b.N; i++ {
		m[i] = Y{}
	}
}

package main

func main() {
	a := 0
	b := 0

	if float64(a)/float64(b) >= 0 || float64(a)/float64(b) <= 0 {
		panic("never")
	}

	if a/b > 0 {
		panic("panics before")
	}
}

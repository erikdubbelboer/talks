package main

func doit(x *int) {
	_ = x
}

func main() {
	y := 1
	doit(&y)
}

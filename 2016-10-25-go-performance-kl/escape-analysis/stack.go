package main

var g *int

func doit(x *int) {
	g = x
}

func main() {
	y := 1
	doit(&y)
}

package main

import (
	"fmt"
)

func f1() {
	var a, b struct{}
	print(&a, "\n", &b, "\n") // Prints same address
	print(&a, "\n", &b, "\n") // Prints same address
	fmt.Println(&a == &b)     // Comparison returns false
}

func f2() {
	var a, b struct{}
	fmt.Printf("%p\n%p\n", &a, &b) // Again, same address
	fmt.Printf("%p\n%p\n", &a, &b) // Again, same address
	fmt.Println(&a == &b)          // ...but the comparison returns true
}

func main() {
	f1()
	f2()
}

// Pointers to distinct zero-size variables may or may not be equal.

// If every pointer to a zero-sized object were required to be different,
// then each allocation of a zero-sized object would have to allocate at least one byte.
// If every pointer to a zero-sized object were required to be the same,
// it would be different to handle taking the address of a zero-sized field within a larger struct.

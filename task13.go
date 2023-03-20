package main

import (
	"fmt"
	"os"
)

func swap() {
	a := 33
	b := 55
	fmt.Fprintln(os.Stdout, a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Fprintln(os.Stdout, a, b)
}

func swap2() {
	a := 13
	b := 15
	fmt.Fprintln(os.Stdout, a, b)
	a, b = b, a
	fmt.Fprintln(os.Stdout, a, b)
}

func task13() {
	swap()
	swap2()
}

package main

import (
	"fmt"
	"os"
)

func intersection[T comparable](setA []T, setB []T) []T {
	var set []T
	values := make(map[T]bool)
	for _, val := range setA {
		values[val] = true
	}
	for _, val := range setB {
		_, ok := values[val]
		if ok {
			set = append(set, val)
		}
	}
	return set
}

func Task11() {
	setA := []int{34, 42, 54, 6, 7}
	setB := []int{8, 5, 4, 6, 23, 42, 54}
	setC := []string{"abs", "add", "mul", "div", "sub", "mod"}
	setD := []string{"adc", "mod", "add", "call"}
	fmt.Fprintln(os.Stdout, intersection(setA, setB))
	fmt.Fprintln(os.Stdout, intersection(setC, setD))
}

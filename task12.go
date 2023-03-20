package main

import (
	"fmt"
	"os"
)

func properSubset[T comparable](setA []T) []T {
	var set []T
	values := make(map[T]struct{})
	for _, val := range setA {
		_, ok := values[val]
		if !ok {
			values[val] = struct{}{}
			set = append(set, val)
		}
	}
	return set
}

func Task12() {
	setA := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Fprintln(os.Stdout, properSubset(setA))
}

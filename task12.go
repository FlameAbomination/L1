package main

import (
	"fmt"
	"os"
)

func properSubset[T comparable](setA []T) []T {
	var set []T
	values := make(map[T]bool)
	for _, val := range setA {
		_, ok := values[val]
		if !ok {
			values[val] = true
			set = append(set, val)
		}
	}
	return set
}

func task12() {
	setA := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Fprintln(os.Stdout, properSubset(setA))
}

package main

import (
	"fmt"
	"os"
	"sync"
)

func Task2() {
	var wg sync.WaitGroup
	arr := [5]int{2, 4, 6, 8, 10}
	for _, val := range arr {
		wg.Add(1)
		go func(v int) {
			// Последовательность вывода не гарантирована
			fmt.Fprintln(os.Stdout, v*v)
			defer wg.Done()
		}(val)
	}

	wg.Wait()
}

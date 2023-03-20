package main

import (
	"fmt"
	"os"
)

func task9() {
	values := make(chan int)
	squares := make(chan int)

	go func() {
		for {
			value := <-values
			squares <- value * value
		}
	}()

	go func() {
		for {
			value := <-squares
			fmt.Fprintln(os.Stdout, value)
		}
	}()

	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		values <- val
	}
}

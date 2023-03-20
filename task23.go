package main

import (
	"errors"
	"fmt"
	"os"
)

func RemoveFromSLice[T any](slice []T, index int) ([]T, error) {
	if index > len(slice) || index < 0 {
		return slice, errors.New("index extends slice length")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

func Task23() {
	arr := []int{2, 4, 6, 8, 10}
	arr, _ = RemoveFromSLice(arr, 4)
	fmt.Fprintln(os.Stdout, arr)
}

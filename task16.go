package main

import (
	"fmt"
	"os"
)

/*
	algorithm quicksort(A, low, high) is
	if low < high then {
		p:= partition(A, low, high)
		quicksort(A, low, p)
		quicksort(A, p + 1, high)
	}
*/

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]
	return i
}

func QuickSort(arr []int) {
	if len(arr) > 1 {
		p := partition(arr)
		QuickSort(arr[:p])
		QuickSort(arr[p+1:])
	}
}

func Task16() {
	arr := []int{10, 8, 4, 2, 6}
	QuickSort(arr)
	fmt.Fprintln(os.Stdout, arr)
}

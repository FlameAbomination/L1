package main

import (
	"fmt"
	"os"
)

func BinarySearch(haystack []int, x int) int {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		mid := (low + high) / 2
		if haystack[mid] < x {
			low = mid + 1
		} else if x < haystack[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func Task17() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Fprintln(os.Stdout, BinarySearch(arr, 8))
}

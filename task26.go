package main

import (
	"fmt"
	"os"
	"strings"
)

func checkString(source string) bool {
	source = strings.ToLower(source)
	values := make(map[rune]bool)
	for _, val := range source {
		_, ok := values[val]
		if ok {
			return false
		} else {
			values[val] = true
		}
	}
	return true
}

func task26() {
	fmt.Fprintln(os.Stdout, checkString("abcd"))
	fmt.Fprintln(os.Stdout, checkString("abCdefAaf"))
	fmt.Fprintln(os.Stdout, checkString("aabcd"))
}

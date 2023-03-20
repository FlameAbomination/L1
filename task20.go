package main

import (
	"fmt"
	"os"
	"strings"
)

// Есть возможность вместо перестановки использовать StringBuilder
func ReverseWords(source string) string {
	sourceSplit := strings.Split(source, " ")
	size := len(sourceSplit) - 1
	for i := 0; i <= size/2; i++ {
		sourceSplit[i], sourceSplit[size-i] = sourceSplit[size-i], sourceSplit[i]
	}
	return strings.Join(sourceSplit, " ")
}

// Есть возможность вместо перестановки использовать StringBuilder
func ReverseWords2(source string) string {
	sourceSplit := strings.Split(source, " ")
	size := len(sourceSplit) - 1
	var builder strings.Builder
	for i := size; i >= 0; i-- {
		builder.WriteString(sourceSplit[i])
		if i != 0 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func Task20() {
	fmt.Fprintln(os.Stdout, ReverseWords("snow главрыба dog sun"))
	fmt.Fprintln(os.Stdout, ReverseWords2("snow главрыба dog sun"))
}

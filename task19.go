package main

import (
	"fmt"
	"os"
	"strings"
)

// Если под тем, что строка подаётся на ходу подразумевается,
// что строка читается слева направо, то решение будет не очень эффективным
func ReverseString(source string) string {
	result := ""
	for _, codepoint := range source {
		result = string(codepoint) + result
	}

	return result
}

// Есть возможность вместо перестановки использовать StringBuilder
func ReverseString2(source string) string {
	runes := []rune(source)
	size := len(runes) - 1
	for i := 0; i <= size/2; i++ {
		runes[i], runes[size-i] = runes[size-i], runes[i]
	}
	return string(runes)
}

// Есть возможность вместо перестановки использовать StringBuilder
func ReverseString3(source string) string {
	runes := []rune(source)
	size := len(runes) - 1
	var builder strings.Builder
	for i := size; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}
func Task19() {
	fmt.Fprintln(os.Stdout, ReverseString("главрыба"))
	fmt.Fprintln(os.Stdout, ReverseString2("главрыба"))
	fmt.Fprintln(os.Stdout, ReverseString3("главрыба"))
}

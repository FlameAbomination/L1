package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func groupTemperature(setA []float64) string {
	var builder strings.Builder
	values := make(map[int][]float64)
	for _, val := range setA {
		valInt := int(val) / 10 * 10
		values[valInt] = append(values[valInt], val)
	}
	for key, value := range values {
		builder.WriteString(strconv.Itoa(key))
		builder.WriteString(":{")
		for i, val := range value {
			builder.WriteString(fmt.Sprintf("%.1f", val))
			if i < len(value)-1 {
				builder.WriteString(", ")
			}
		}
		builder.WriteString("}, ")
	}
	return builder.String()[:builder.Len()-2]
}

func Task10() {
	setA := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Fprintln(os.Stdout, groupTemperature(setA))
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func SetBit(value int64, bit int8, state uint) (int64, error) {
	if state > 1 {
		return int64(value), errors.New("state must be 0 or 1")
	}
	if bit > 63 {
		return int64(value), errors.New("bit index cannot extend 64 bits")
	}
	if state == 0 {
		// Если бит нужно очистить, то использует and с инвертированной маской
		value &= ^(1 << bit)
	} else {
		// Если бит нужно поставить в 1, то использует or с маской
		value |= (1 << bit)
	}
	return value, nil
}

func task8() {
	var value int64
	value = 42
	fmt.Fprintln(os.Stdout, strconv.FormatInt(value, 2))
	value, _ = SetBit(value, 4, 1)
	fmt.Fprintln(os.Stdout, strconv.FormatInt(value, 2))
}

package main

import (
	"fmt"
	"math/big"
	"os"
)

func BigIntAdd(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

func BigIntSub(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

func BigIntDiv(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

func BigIntMul(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func task22() {
	a := new(big.Int)
	b := new(big.Int)
	a.SetInt64(2 << 21)
	b.SetInt64(2 << 22)
	fmt.Fprintln(os.Stdout, BigIntAdd(a, b))
	fmt.Fprintln(os.Stdout, BigIntSub(a, b))
	fmt.Fprintln(os.Stdout, BigIntDiv(a, b))
	fmt.Fprintln(os.Stdout, BigIntMul(a, b))
}

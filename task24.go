package main

import (
	"fmt"
	"math"
	"os"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	point := Point{
		x: x,
		y: y,
	}
	return point
}

func (point Point) Distance(point2 Point) float64 {
	return float64(math.Sqrt(math.Pow(point.x-point2.x, 2) + math.Pow(point.y-point2.y, 2)))
}

func Task24() {
	a := NewPoint(1, 1)
	b := NewPoint(0, 0)
	fmt.Fprintln(os.Stdout, a.Distance(b))
}

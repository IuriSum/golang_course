package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height, width float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{10, 10}
	s := square{10}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.height * t.width / 2
}

func (s square) getArea() float64 {
	return math.Pow(s.side, 2.0)
}

func printArea(s shape) {
	fmt.Println("Area is", s.getArea())
}

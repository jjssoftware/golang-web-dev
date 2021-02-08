package main

import (
	"fmt"
	"math"
)

type square struct {
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.width
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{ 4.25 }
	c := circle{ 2.3 }

	info(s)
	info(c)
}

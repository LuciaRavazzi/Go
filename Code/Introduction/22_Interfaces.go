package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	height float64
	width  float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// If a struct contains area, it will be executed.
type shape interface {
	area() float64
}

func main() {

	c1 := circle{4.5}
	r1 := rect{5, 7}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}

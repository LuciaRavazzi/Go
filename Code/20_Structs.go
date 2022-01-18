package main

import "fmt"

// Structs are the class type in Go.
// Custom type.
type Point struct {
	x int32
	y int32
}

func changeX(pt *Point) {
	pt.x = 100
}

type Circle struct {
	radius float64
	center *Point
}

func main() {

	var p1 Point = Point{1, 2}
	fmt.Println(p1.x, p1.y)

	p2 := Point{x: 1, y: 3}
	fmt.Println(p2)

	p3 := &Point{x: 1, y: 3}
	fmt.Println(p3)
	changeX(p3)
	fmt.Println(p3)

	c1 := Circle{4.56, &p1}
	fmt.Println(c1.center.x, c1.center.y, c1.radius)
}

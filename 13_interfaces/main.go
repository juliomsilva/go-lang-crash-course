package main

import (
	"fmt"
	"math"
)

// Shapre
type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

// Circle
type circle struct {
	x, y, radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Rectangle
type rectangle struct {
	width, heigth float64
}

func (r rectangle) area() float64 {
	return r.width * r.heigth
}

func main() {
	circle := circle{x: 0, y: 0, radius: 5}
	rectangle := rectangle{width: 10, heigth: 5}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectangle))
}

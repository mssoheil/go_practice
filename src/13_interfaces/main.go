package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	with, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.with * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {

	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{with: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}

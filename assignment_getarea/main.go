package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	width  float64
	height float64
}

func main() {
	shapes := []shape{
		&triangle{base: 5, height: 10},
		&square{width: 5, height: 5},
	}

	for _, shape := range shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.getArea())
	}
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.width * s.height
}

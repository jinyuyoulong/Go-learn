package main

import (
	"math"
)

type Rectangle struct {
	width  float64
	height float64
}

type Cycle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return rectangle.width*2 + rectangle.height*2
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Cycle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

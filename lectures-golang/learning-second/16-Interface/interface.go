package main

import (
	"fmt"
	"math"
)

// interface contains a method called area

type form interface {
	area() float64
}

// A function is created that receives a parameter related to the interface.
// In the example, it receives the "form" interface

func writeArea(f form) {
	fmt.Printf("The area of the form is %0.2f\n", f.area())
}

type rectangle struct {
	height float64
	width  float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	ray float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.ray, 2)
}

func main() {
	r := rectangle{10, 15}
	writeArea(r)

	c := circle{5}
	writeArea(c)
}

package main

import (
	"fmt"
	"math"
)

type form interface {
	area() float64
}

type rectangle struct {
	height float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func writeArea(f form) {
	fmt.Println(f.area())
}

func genericInterface(interf interface{}) { // act as a generic type
	fmt.Println(interf)
}

func main() {
	f1 := rectangle{height: 5, width: 5}
	f2 := circle{radius: 5}
	writeArea(f1)
	writeArea(f2)

	genericInterface(f1)
	genericInterface(f2)
	genericInterface("string")
	genericInterface(-10)
}

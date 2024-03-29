package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Squaer struct {
	size float64
}

func (s Squaer) Area() float64 {
	return s.size * s.size
}

func (s Squaer) Perimeter() float64 {
	return s.size * 4
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

func main() {
	// var s Shape = Squaer{3}
	// fmt.Printf("%T\n", s)
	// fmt.Println("Area: ", s.Area())
	// fmt.Println("Perimeter: ", s.Perimeter())

	var s Shape = Squaer{3}
	printInformation(s)

	c := Circle{6}
	printInformation(c)
}

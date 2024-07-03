package main

import (
	"fmt"
	"math"
)

// interface and type implementations
type Shape interface {
	Area() float64
	Permimeter() float64
}

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

// type needs to provide the definition for all the methods declared in the interface
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// 2. Implementing the Shape interface for Rectangle

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Permimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 1. Implementing the Shape interface for Rectangle

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Permimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Here, Circle implements the Shape interface by defining both Area() and Perimeter() methods.

// 3. Interface polymorphism

// interfaces in go enable polymorphism. allowing different types to used interchangeably. based on the share methods
// this means we can write functions or methods that accepts interfaces as parameters
// and operate on any type that satisfies the interface

// The PrintShapeDetails function can accept any type that implements the Shape interface. For example:
func PrintShapeDetails(s Shape) {
	fmt.Println("Area: ", s.Area())
	fmt.Println("Permimeter: ", s.Permimeter())
}

func PrintAnimalSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	c := Circle{Radius: 3}
	PrintShapeDetails(c)

	d := Rectangle{
		Width:  3,
		Height: 4,
	}
	PrintShapeDetails(d)

	dog := Dog{}
	PrintAnimalSound(dog)
}

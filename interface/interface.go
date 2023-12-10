package main

import "fmt"

// Let's define an interface named 'Shape' with a single method 'Area() float64'
type Shape interface {
	Area() float64
}

// Rectangle struct implementing the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct implementing the Shape interface
type Circle struct {
	Radius float64
}

// Implementing the Area() method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implementing the Area() method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape

	s = Rectangle{Width: 5, Height: 3}
	fmt.Println("Area of Rectangle:", s.Area()) // Output: Area of Rectangle: 15

	s = Circle{Radius: 2}
	fmt.Println("Area of Circle:", s.Area()) // Output: Area of Circle: 12.56
}

// improve this code

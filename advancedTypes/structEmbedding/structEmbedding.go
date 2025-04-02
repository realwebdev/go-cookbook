package main

import "fmt"

// stringer method

type Color int

func (c Color) String() string { // fmt stringer
	switch c {
	case ColorBlack:
		return "Black"
	case ColorBlue:
		return "Blue"
	case ColorYellow:
		return "Yellow"
	case ColorPink:
		return "Pink"
	default:
		panic("invalid color given")
	}
}

const (
	ColorBlue Color = iota
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {
	fmt.Printf("the color is %v", ColorPink)
}

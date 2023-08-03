package main

import "fmt"

type Person struct {
	Name string
}

type Hamza struct {
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}

func (h Hamza) SayGoodbye() {
	// Check if h is a Person struct.
	if h, ok := h.(Person); ok {
		// Cast h to a Person struct and call the SayHello() method on it.
		h.SayHello()
	}
}

func main() {
	p := Person{Name: "John Doe"}

	// Call the SayHello() method on p.
	p.SayHello()

	// Call the SayGoodbye() method on p.
	h := Hamza{}
	h.SayGoodbye()
}

package main

import "fmt"

// struct embedding

type Entity struct {
	Name    string
	Version float32
	ID      int
	PosX    float32
	PosY    float32
}

type SpecialEntity struct {
	Entity
	SpecialField float32
}

func main() {
	e := SpecialEntity{
		SpecialField: 88.8,
		Entity: Entity{
			Name: "special",
			Version: 2,
			ID: 221,
			PosX: 22.00,
			PosY: 11.3,
		},
	}

	fmt.Printf("%+v\n", e)
}

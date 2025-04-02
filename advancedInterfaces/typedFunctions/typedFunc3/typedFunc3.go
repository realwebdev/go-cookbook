package main

import (
	"fmt"
	"strings"
)

// function composition
type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(s string) string {
	return "FOO_" + s
}

func compose(f1, f2 TransformFunc) TransformFunc {
	return func(s string) string {
		return f2(f1(s))
	}
}

// function to compose two transformations
func main() {
	combined := compose(Uppercase, Prefixer)

	fmt.Println(combined("hello"))

}

package main

import (
	"fmt"
	"strings"
)

type TransFormFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// a function that returns a prefixing with a dynamic prefix
func Prefixer(prefix string) TransFormFunc {
	return func(s string) string {
		return prefix + s
	}
}

// function to compose

func compose(funcs ...TransFormFunc) TransFormFunc {
	return func(s string) string {
		for _, fn := range funcs {
			s = fn(s)
		}
		return s
	}
}

func main() {
	// Use Prefixer to create different functions dynamically
	fooPrefix := Prefixer("FOO_")
	barPrefix := Prefixer("BAR_")

	composedFunc := compose(Uppercase, fooPrefix) // ✅ Step 1: Returns a function
	result := composedFunc("hello")               // ✅ Step 2: Call the returned function
	fmt.Println(result)
	fmt.Println(compose(Uppercase, barPrefix)("hello"))
}

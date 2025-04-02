// functional programming style

package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func Prefixer(s string) string {
	return "FOO_" + s
}

// higher order function that applier a transformation to each string in a slice
func mapStrings(strings []string, fn TransformFunc) []string {
	result := make([]string, len(strings))
	for i, s := range strings {
		result[i] = fn(s)
	}

	return result
}

func main() {
	words := []string{"hello", "world!"}

	// apply transformations dynamically
	fmt.Println(mapStrings(words, Uppercase))
	fmt.Println(mapStrings(words, Prefixer))
}

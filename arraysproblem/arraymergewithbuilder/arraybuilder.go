package main

import (
	"fmt"
	"strings"
)

func main() {
	array1 := "abc"
	array2 := "pqrstu"

	result := mergeAlternatively(array1, array2)
	fmt.Println(result)
}

func mergeAlternatively(word1 string, word2 string) string {
	var result strings.Builder

	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result.WriteByte(word1[i])
		result.WriteByte(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result.WriteByte(word1[i])
		i++
	}

	for j < len(word2) {
		result.WriteByte(word2[j])
		j++
	}

	return result.String()
}

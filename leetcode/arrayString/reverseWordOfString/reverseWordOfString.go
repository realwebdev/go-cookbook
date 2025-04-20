package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// trim the leading and trailing spaces
	trimmed := strings.TrimSpace(s)

	// split the string into words using whitespace as a delimiter
	words := strings.Fields(trimmed)

	// reverse the order of the words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// join the reversed string with the single space
	return strings.Join(words, " ")

}

func main() {
	words := "a good  example"
	fmt.Println(reverseWords(words))

}

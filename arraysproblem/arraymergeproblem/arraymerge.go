// array merge problem

package main

import "fmt"

func main() {
	array1 := "abc"
	array2 := "pqrstu"

	result := mergeAlternately(array1, array2)
	fmt.Println(result)
}

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result += string(word1[i]) + string(word2[j])
		i++
		j++
	}

	for i < len(word1) {
		result += string(word1[i])
		i++
	}

	for j < len(word2) {
		result += string(word2[j])
		j++
	}

	return result
}

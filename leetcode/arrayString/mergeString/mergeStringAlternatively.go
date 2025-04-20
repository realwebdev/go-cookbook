package main

import "fmt"

func mergeAlternatively(word1 string, word2 string) string {
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

	}

	return result

}

func main() {
	w1 := "abcd"
	w2 := "pqre"
	ans := mergeAlternatively(w1, w2)
	fmt.Println(ans)
}

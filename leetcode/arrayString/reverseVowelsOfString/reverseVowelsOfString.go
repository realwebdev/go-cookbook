package main

import "fmt"

func reverseVowels(s string) string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	runes := []rune(s)
	left := 0
	right := len(runes) - 1

	for left < right {
		// move left pointer until it points a vowel
		for left < right && !vowels[runes[left]] {
			left++
		}

		// move right pointer until it points to a vowel
		for left < right && !vowels[runes[right]] {
			right--
		}

		// swap the vowels
		if left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}
	}

	return string(runes)

}

func main() {
	s1 := "hello"
	fmt.Println(reverseVowels(s1))

	s2 := "leetcode"
	fmt.Println(reverseVowels(s2))

}

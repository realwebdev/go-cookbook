package main

import (
	"fmt"
	"sort"
)

// concept -> same set of character, same frequency multiset
// solution
// 1. count frequency on each characters in both words
// 2. ensure both have same unique characters
// 3. ensure both words have same frequency counts. Regardless of which letter has which frequency.

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	// frequency count maps
	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)

	// character set
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)

	for _, ch := range word1 {
		freq1[ch]++
		set1[ch] = true
	}

	for _, ch := range word2 {
		freq2[ch]++
		set2[ch] = true
	}

	// compare the characters
	if len(set1) != len(set2) {
		return false
	}

	for ch := range set1 {
		fmt.Println(set1[ch])
		fmt.Println(set2[ch])
		if !set2[ch] {
			return false
		}
	}

	// compare frequency values regardless of keys

	values1 := []int{}
	values2 := []int{}

	for _, v := range freq1 {
		values1 = append(values1, v)
	}
	for _, v := range freq2 {
		values2 = append(values2, v)
	}

	sort.Ints(values1)
	sort.Ints(values2)

	for i := range values1 {
		if values1[i] != values2[i] {
			return false
		}
	}

	return true
}

func main() {
	arr1 := "aabbcc"
	arr2 := "xxyyzz"
	t := closeStrings(arr1, arr2)
	fmt.Println("this is the result", t)
}

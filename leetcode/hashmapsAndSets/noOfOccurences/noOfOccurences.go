package main

import "fmt"

func uniqueOccurrences(arr []int) bool {
	freqMap := make(map[int]int)
	countMap := make(map[int]bool) // store no of times

	// counting the frequency of each number
	for _, num := range arr {
		freqMap[num]++
	}

	// check if the frequencies are unique
	for _, freq := range freqMap {
		if countMap[freq] {
			return false // dublicate freq found
		}
		countMap[freq] = true
	}

	return true
}

func main() {
	arr := []int{1, 2, 2, 1, 1, 3}
	fmt.Println(uniqueOccurrences(arr))

	arr2 := []int{1, 2, 2, 1, 1, 3, 3, 3}
	fmt.Println(uniqueOccurrences(arr2))

}

package main

import (
	"fmt"
	"math"
)

func main() {
	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}
	answer := findMaxAverageSlidingWindow(nums, k)
	fmt.Println(answer)
}

func findMaxAverageSlidingWindow(nums []int, k int) float64 {
	var windowSum int
	var start int
	max := math.Inf(-1)

	for end := 0; end < len(nums); end++ {
		windowSum += nums[end]

		if (end - start + 1) == k {
			// calculate the average
			max = math.Max(max, float64(windowSum)/float64(k))
			windowSum -= nums[start]
			start += 1
		}
	}
	return max
}

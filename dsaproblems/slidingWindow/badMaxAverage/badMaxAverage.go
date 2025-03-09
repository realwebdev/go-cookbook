package main

import (
	"fmt"
	"math"
)

func main(){
	k := 4
	nums := []int{1, 12, -5, -6, 50, 3}
	answer := findMaxAverage(nums,k)
	fmt.Println(answer)
}

func findMaxAverage(nums []int,k int) float64{
	max := math.Inf(-1)
	for i := 0; i < len(nums); i++ {
		windowSum := 0
		for j := 0; j < k; j++ {
			if !(i+k-1 > len(nums)-1) {
				windowSum += nums[i+j]
			} else {
				i = len(nums)
				j = k
			}
		}
		if windowSum != 0 {
			max = math.Max(max, float64(windowSum)/float64(k))
		}
	}
	if max == math.Inf(-1) {
		return 0
	}

	return max
	

}
package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// calculate the left products
	leftProduct := 1
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		leftProduct *= nums[i]
	}

	// calculate the right products and final result
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]

	}

	return answer

}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // Output: [24 12 8 6]

	nums2 := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums2)) // Output: [0 0 9 0 0]

}

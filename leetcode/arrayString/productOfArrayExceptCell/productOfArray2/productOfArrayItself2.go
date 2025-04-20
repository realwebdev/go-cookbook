package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	leftProduct := 1
	fmt.Println("== LEFT PASS ==")
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		fmt.Printf("i=%d | nums[i]=%d | leftProduct(before)=%d | answer[%d]=%d\n",
			i, nums[i], leftProduct, i, answer[i])
		leftProduct *= nums[i]
	}

	fmt.Println("\n== RIGHT PASS ==")
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("i=%d | nums[i]=%d | rightProduct(before)=%d | answer[%d]=%d * %d = ",
			i, nums[i], rightProduct, i, answer[i], rightProduct)
		answer[i] *= rightProduct
		fmt.Printf("%d\n", answer[i])
		rightProduct *= nums[i]
	}

	fmt.Println("\n== FINAL RESULT ==")
	fmt.Println("answer:", answer)
	return answer
}

func main() {
	nums := []int{1, 2, 3, 4}
	productExceptSelf(nums)
}

// package main

// import "fmt"

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	answer := make([]int, n)

// 	// Step 1: Fill left products
// 	answer[0] = 1
// 	for i := 1; i < n; i++ {
// 		answer[i] = nums[i-1] * answer[i-1]
// 	}

// 	// Step 2: Multiply with right products on the fly
// 	right := 1
// 	for i := n - 1; i >= 0; i-- {
// 		answer[i] *= right
// 		right *= nums[i]
// 	}

// 	return answer

// }

// func main() {
// 	nums1 := []int{1, 2, 3, 4}
// 	fmt.Println(productExceptSelf(nums1)) // Output: [24 12 8 6]

// 	nums2 := []int{-1, 1, 0, -3, 3}
// 	fmt.Println(productExceptSelf(nums2)) // Output: [0 0 9 0 0]

// }

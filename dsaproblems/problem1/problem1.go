package main

import "fmt"

/*
### **🔥 Problem 1: Sliding Window Maximum**
📌 **Problem Statement:**
Given an array `nums` and an integer `k`, find the maximum value in each contiguous subarray of size `k`.

📌 **Example:**
```go
Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
```
📌 **Explanation:**
- Window `[1,3,-1]` → max = `3`
- Window `[3,-1,-3]` → max = `3`
- Window `[-1,-3,5]` → max = `5`
- And so on...

📌 **Constraints:**
- `1 <= nums.length <= 10^5`
- `1 <= k <= nums.length`
- `nums[i]` can be **negative or positive**

### **⏳ Your Task:**
1. Solve this problem in **Go**.
2. Optimize for **time complexity** (brute force is `O(n*k)`, try for `O(n)`).
3. Use **Golang slices, deque, or heap** (if needed).

Give it a shot! 💪 Let me know once you've tried, and I'll help improve your solution. 🚀
*/
func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	result := maxSlidingWindow(nums, k)
	fmt.Println(result) // Output: [3,3,5,5,6,7]
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	// Initialize the result slice
	result := make([]int, 0)

	// Initialize the deque
	deque := make([]int, 0)

	// Iterate over the nums slice
	for i, num := range nums {
		// Remove the first element if it's out of the window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove the last element if it's smaller than the current element
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}

		// Add the current element to the deque
		deque = append(deque, i)

		// Add the maximum element to the result slice
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

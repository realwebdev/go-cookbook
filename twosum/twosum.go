package main

import "fmt"

// TwoSum finds the indices of the two numbers that add up to the target
func TwoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if index, ok := numMap[complement]; ok {
            return []int{index, i}
        }
        numMap[num] = i
    }
    return nil
}

func main() {
    // Example usage
    nums := []int{2, 7, 11, 15}
    target := 9
    result := TwoSum(nums, target)
    fmt.Println(result) // Output: [0, 1]
}

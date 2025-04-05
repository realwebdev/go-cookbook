package main

import "fmt"

// find the difference of two arrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// storing unique elements in sets
	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}

	// find element in nums 1 but not in nums 2
	res1 := []int{}
	for num := range set1 {
		fmt.Println(set2[num])
		if !set2[num] {
			res1 = append(res1, num)
		}
	}

	// find element in nums 2 but not in nums 1
	res2 := []int{}
	for num := range set2 {
		if !set1[num] {
			res2 = append(res2, num)
		}
	}

	return [][]int{res1, res2}

}

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	result := findDifference(nums1, nums2)
	fmt.Println(result) // [[1,3], [4,6]]
}

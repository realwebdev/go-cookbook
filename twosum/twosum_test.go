package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2, 3}, 7, nil}, // Test case where no solution exists
	}

	for _, tc := range testCases {
		got := TwoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("TwoSum(%v, %d) = %v; want %v", tc.nums, tc.target, got, tc.want)
		}
	}
}

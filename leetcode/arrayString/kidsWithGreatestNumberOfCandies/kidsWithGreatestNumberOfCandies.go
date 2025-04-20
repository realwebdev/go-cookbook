package main

import "fmt"

/*
Here's the step-by-step approach:

1. Find the Maximum: First, we need to find the maximum number of candies
any kid has in the original candies array. Let's call this max_candies.

2. Iterate and Check: Then, we iterate through the candies array again
(or we can do this in a single pass if we keep track of the maximum).
For each kid i with candies[i] candies, we check if candies[i] + extraCandies >= max_candies.

3. Build the Result Array: Based on the comparison in step 2, we build a boolean array result.
If the condition candies[i] + extraCandies >= max_candies is true, we set result[i] to true;
otherwise, we set it to false.

4. Return the Result: Finally, we return the result array.
*/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, candyCount := range candies {
		if candyCount > maxCandies {
			maxCandies = candyCount
		}
	}

	result := make([]bool, len(candies)) // boolean array size of candies
	for i, candyCount := range candies {
		if candyCount+extraCandies >= maxCandies {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result

}

func main() {
	candies1 := []int{2, 3, 4, 5, 6, 7}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies1, extraCandies))

}

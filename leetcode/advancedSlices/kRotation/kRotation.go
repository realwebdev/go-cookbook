package main

import "fmt"

func reverse(){
	
}

func kShift(a []int) []int {

	// last := a[len(a)-1]

	// hardcoded try
	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }

	// for i, j := 0, 2; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }

	// for i, j := 3, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }

	return a

}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	shiftedByK := kShift(nums)

	fmt.Printf("shifted by K --> %d", shiftedByK)
}

package main

import (
	"fmt"
)

// Function that removes duplicates from a sorted array
// returns the number of unique elements in the 'nums' array passed in
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
			return 0
	}

	k := 1 // Initialize k to 1, as the first element is always unique

	for i := 1; i < len(nums); i++ {
		// If the current element is different from the previous element,
		// move it to the k-th position in the array.
		if nums[i] != nums[i - 1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
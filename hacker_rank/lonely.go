package main

import (
	"fmt"
	"sort"
)

// function that checks for a number that occurs only once in an
// array
func lonelyInteger(a []int32) int32 {
	// sort the slice first
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j]})

	numbers := make(map[int32]int)

	for _, val := range a {
		numbers[val]++
	}

	for val, occ := range numbers {
		if occ == 1 {
			return val
		}
	}

	return 0
}

func main() {
	fmt.Println(lonelyInteger([]int32{1,2,3,4,3,2,1}), "<= Expecting 4")
}
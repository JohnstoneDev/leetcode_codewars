/*
	Given two sorted arrays nums1 and nums2 of size m and n respectively,
	return the median of the two sorted arrays.
*/

package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrs(nums1, nums2 []int) float64 {
	mergedSlice := append(nums1, nums2...)
	var median float64

	// sort the merged slice
	sort.Ints(mergedSlice)

	if len(mergedSlice) % 2 == 0 {
		x := len(mergedSlice) / 2
		y := x - 1
		median = float64(mergedSlice[x] + mergedSlice[y]) / 2
	} else {
		middleIndex := len(mergedSlice) / 2
		median = float64(mergedSlice[middleIndex])
	}

	return median
}

func main() {
	fmt.Println(findMedianSortedArrs([]int{1, 2}, []int{3, 4}), "<= Expect 2")
	fmt.Println(findMedianSortedArrs([]int{1, 3}, []int{2}), "<= Expect 2")
}
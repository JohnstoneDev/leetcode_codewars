package main

import (
	"fmt"
)

/*
	Given a divisor & a Bound. Find the largest integer N
	such that :

	N is divisible by divisor
	N <= bound
	n > 0
*/

func MaxMultiple(divisor, bound int) int {
	numbers := []int{}

	for i := 1; i <= bound; i++ {
		if i % divisor == 0 {
			numbers = append(numbers, i)
		}
	}

	return numbers[len(numbers) - 1]
}

func main() {
	fmt.Println(MaxMultiple(2,7), "<= Got, Expected 6")
	fmt.Println(MaxMultiple(37, 200), "<= Got, Expected 185")
	fmt.Println(MaxMultiple(7, 17), "<= Got, Expected 14")
}
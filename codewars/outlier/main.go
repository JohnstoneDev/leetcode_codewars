package main

import (
	"fmt"
)

func ArrayCondition(intergers []int) string {
	countOdd := 0
	countEven := 0

	for _, num := range intergers {
		if num % 2 == 0 {
			 countEven++
			} else {
				countOdd++
			}
	}

	if countOdd > countEven {
		return "odd"
	}

	return "even"
}

func FindOutlier(integers []int) int {
	condition := ArrayCondition(integers)

	fmt.Println(condition)

	if condition == "odd" {
		for _, num := range integers {
			if num % 2 == 0 {
				return num
			}
		}
	} else {
		for _, num := range integers {
			if num % 2 != 0 {
				return num
			}
		}
	}

	return 0
}

func main(){
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}
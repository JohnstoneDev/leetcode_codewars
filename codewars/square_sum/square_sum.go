/*
	Square(n)Sum : Squares each number passed into it &
	sums the results together
*/

package main

import (
	"fmt"
)

func SquareSum(numbers []int) int {
	var sum int;

	for _, num := range numbers {
		sum += (num * num)
	}

	return sum
}

func main(){
	fmt.Println(SquareSum([]int{0,3,4,5}))
}
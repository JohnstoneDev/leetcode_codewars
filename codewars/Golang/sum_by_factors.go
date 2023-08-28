package main

import (
	"fmt"
)

func SumOfDivided(ls []int) string {
	var sorted []int;

	fmt.Println(sorted)

	return "Hello there!"
}


func main() {
	fmt.Println(SumOfDivided([]int{12, 15}), "<= Got,  Expecting:", "(2 12)(3 27)(5 15)")
	fmt.Println("----------------------------------------------------------------------------------------")
	fmt.Println(SumOfDivided([]int{15,21,24,30,45}), "<= Got,  Expecting:", "(2 54)(3 135)(5 90)(7 21)")
	fmt.Println("----------------------------------------------------------------------------------------")
}
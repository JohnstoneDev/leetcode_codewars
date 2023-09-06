package main

import (
	"fmt"
	"strconv"
)

/*
	Write a function that accepts an array of 10 integers (between 0 and 9),
	that returns a string of those numbers in the form of a phone number.
*/

func CreatePhone(numbers [10]uint) string {
	phone := ""

	for _, num := range numbers {
		converted := strconv.Itoa(int(num))
		phone += converted
	}

	nums := []rune{'('}

	for i, char := range phone {
		if i <= 2 {
			nums = append(nums, char)
		}
	}

	nums = append(nums, ')', ' ')

	for i, char := range phone {
		if i > 2 && i <= 5 {
			nums = append(nums, char)
		}
	}

	nums = append(nums, '-')

	for i, char := range phone {
		if i >= 6 {
			nums = append(nums, char)
		}
	}

	return string(nums)
}


func CreatePhoneNumber(numbers [10]uint) string {
	phone := fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d",
		numbers[0], numbers[1], numbers[2],
		numbers[3], numbers[4], numbers[5],
		numbers[6], numbers[7], numbers[8], numbers[9])

	return phone
}

func main(){
	fmt.Println(CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0}), "<= returns (123) 456-7890")  // returns "(123) 456-7890"
	fmt.Println(CreatePhone([10]uint{1,2,3,4,5,6,7,8,9,0}), "<= returns (123) 456-7890")  // returns "(123) 456-7890"
}
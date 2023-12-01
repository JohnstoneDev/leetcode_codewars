package main

import (
	"fmt"
	"strings"
)

// capitalize first letter of every word "Jaden Case"
func ToJadenCase(str string) string {
	words := strings.Split(str, " ")

	for i, word := range words {
		letters := strings.Split(word, "")
		letters[0] = strings.ToUpper(letters[0])
		words[i] = strings.Join(letters, "")
	}

	return strings.Join(words, " ")
}

// using the standard library
// func ToJadenCase(str string) string {
// 	return strings.Title(str)
// }


func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	fmt.Println(ToJadenCase("When I die. then you will realize"))
	fmt.Println(ToJadenCase("Jonah Hill is a genius"))
	fmt.Println(ToJadenCase("Dying is mainstream"))
}

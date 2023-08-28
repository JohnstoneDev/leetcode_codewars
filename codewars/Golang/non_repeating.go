/*
	Program that takes a string input & returns the first character that
	is not repeated anywhere in the string, Case insensitive while checking
	Case sensitive while returning
*/

package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	// if the string is empty return an empty string
	if len(str) <= 0 {
		return ""
	}

	unique := []rune{} // slice that stores unique letters
	downcased := strings.ToLower(str) // downcase the string for comparison
	characters := make(map[rune]int) // map that stores occurence of a character in the string

	// load the map
	for _, letter := range downcased {
		characters[letter] += 1
	}

	for i, letter := range downcased {
		if characters[letter] == 1 {
			unique = append(unique, rune(str[i]))
		}
	}

	if len(unique) <= 0 {
		return ""
	}

	return string(unique[0])
}

func main(){
	fmt.Println(FirstNonRepeating("abba"), "<= Got, Expected: ")
	fmt.Println(FirstNonRepeating("~><#~><"), "<= Got, Expected: #")
	fmt.Println(FirstNonRepeating("sTreSS"), "<= Got, Expected: T")
	fmt.Println(FirstNonRepeating("hello world, eh?"), "<= Got, Expected: w")
}
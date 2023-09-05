package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Split(str, " ")

	for index , word := range words {
		reversed := []rune(word)

		for i, j := 0, len(reversed) - 1; i < j; i, j = i + 1, j - 1 {
			reversed[i], reversed[j] =  reversed[j], reversed[i]
		}

		word = string(reversed)
		words[index] = word
	}

	str = strings.Join(words, " ")

	return str
}

func main() {
	fmt.Println(ReverseWords("The quick brown fox jumps over the lazy dog."), "<= Got, Expecting: ehT kciuq nworb xof spmuj revo eht yzal .god")
}
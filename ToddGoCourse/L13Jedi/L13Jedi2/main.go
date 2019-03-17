package main

import (
	"fmt"
	"goprogramming/ToddGoCourse/L13Jedi/L13Jedi2/quote"
	"goprogramming/ToddGoCourse/L13Jedi/L13Jedi2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}

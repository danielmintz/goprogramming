// Package word is all about the magic of words
package word

import "strings"

// UseCount returns number of times a word appears
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count adds up all the words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}

// write the code for this func

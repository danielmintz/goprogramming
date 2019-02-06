package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{6, 2, 8, 4, 9, 3, 5, 22, 1, 3, 44, 77, 88}
	xs := []string{"Daniel", "Marni", "Debra", "Josef", "Scrivener", "Mintz", "Humphries", "Bursik"}

	fmt.Println(xi)
	sort.Ints(xi) // sorts it in order
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

//The sort package allows us to sort slices.

package main

import (
	"fmt"
)

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 9}
	i := sum(ii...)
	fmt.Println(i)

	s1 := bar(sum)
	fmt.Println(s1)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total

}

func even (f func(xi ...int) int, {
	return f

}

//Hands-on exercise #9
//A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument

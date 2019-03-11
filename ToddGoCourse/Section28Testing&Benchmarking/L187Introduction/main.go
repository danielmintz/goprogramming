package main

import "fmt"

func main() {
	fmt.Println(mySum(2, 3))
	fmt.Println(mySum(6, 7, 8, 9))
}

func mySum(xs ...int) int {
	sum := 0
	for _, v := range xs {
		sum += v

	}
	return sum

}

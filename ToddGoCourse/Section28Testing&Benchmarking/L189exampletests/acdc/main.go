// Package acdc is all about everything rock and roll
package acdc

// Sum adds up all the numbers
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum

}

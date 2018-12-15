package main

import "fmt"

func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
func main() {
	fmt.Println(max(100, 15) + max(14, 15))
}

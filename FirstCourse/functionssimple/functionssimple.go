package main

import "fmt"

func max(i int, j int) {
	if i > j {
		fmt.Println(i)
	} else {
		fmt.Println(j)
	}
}
func main() {
	max(10, 15)
}

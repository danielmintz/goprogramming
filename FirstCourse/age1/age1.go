package main

import "fmt"

func main() {
	age := 10

	if age < 13 {
		fmt.Println("Wow, you are young!")
	} else if age < 20 {
		fmt.Println("You're a teenager)")
	} else if age < 30 {
		fmt.Println("You're in your twenties")
	} else if age < 40 {
		fmt.Println("You're in your thirties")
	} else if age < 50 {
		fmt.Println("You're getting there")
	} else {
		fmt.Println("over the hill")
	}

}

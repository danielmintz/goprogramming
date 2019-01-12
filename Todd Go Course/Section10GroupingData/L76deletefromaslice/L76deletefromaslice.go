package main

import "fmt"

func main() {

	x := []int{2, 4, 5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])
	fmt.Println(x[3:])
	fmt.Println(x[1:4])

	for i, v := range x {
		fmt.Println(i, v)

	}
	y := []int{56, 77, 88, 99}
	fmt.Println(y)
	x = append(x, 566, 777)
	fmt.Println(x)
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[0:4], x[5:8]...)
	fmt.Println(x)
}

//We can delete from a slice using both append and slicing. This is a wonderful and elegant example
//of why Go is great and how Go provides ease of programming.

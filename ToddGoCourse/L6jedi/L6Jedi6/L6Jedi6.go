package main

import "fmt"

func main() {

	func() {
		fmt.Println("hello anonymous func")
	}()

	func(x int) {
		fmt.Println("hello this is an anonymous func with a variabe", x)
	}(24)

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

	}()
	fmt.Println("Done")

	func(i int) {
		for ; i < 19; i++ {
			fmt.Println(i)
		}

	}(1)
	fmt.Println("Done")
}

//Hands-on exercise #6
//Build and use an anonymous func

//func (r receiver) identifier(parameters) (returns) { code }

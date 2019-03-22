package main

import "fmt"

func main() {

	message := "Hello"
	sayhello(&message) // memory address
	println(message)
}

func sayhello(message *string) { //pointer to a memeory address
	fmt.Println(*message) //dereferencing to return value

	*message = "Hello Go"
}

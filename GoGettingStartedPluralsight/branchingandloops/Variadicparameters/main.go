package main

import "fmt"

func main() {

	sayhello("Hello", "gg", "from", "Plurlaisght")
}

func sayhello(messages ...string) {
	for _, v := range messages {
		fmt.Println(v)
	}
}

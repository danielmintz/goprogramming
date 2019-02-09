package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello Playground") // The above is relaity in the background is just calling the below 
												//(Methods, type, interfaces)
	io.WriteString(os.Stdout, "Hello Playground")

}

//Interface says: If you got these methods then you are my type

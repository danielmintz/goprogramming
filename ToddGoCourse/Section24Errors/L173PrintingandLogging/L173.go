package main

import (
	"fmt"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.text")
	if err != nil {
		// fmt.Println("err happened ", err)
		// log.Println("err happened ", err)
		// log.Fatalln(err)
		// log.Panicln(err)
		panic(err)
	}
}

func foo() {
	fmt.Println("when os.Exit() is called defered functions don't rund ")
}

//... the Fatal functions call os.Exit(1) after writing the log message ...
// Fatalln is equivalent to Println() followed by a call to os.Exit(1).

// Printing & logging
// You have a few options to choose from when it comes to printing out, or logging, an error message:
// fmt.Println()
// log.Println()
// log.Fatalln()
// os.Exit()
// log.Panicln()  --- deferred functions run
// can use “recover”
// panic()

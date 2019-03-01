package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("Log_File.text")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.text")
	if err != nil {
		//fmt.Println("err happened ", err)
		log.Println("err happened ", err)
	}
	defer f2.Close()

	fmt.Println("check the 'log_File.text in the directory")
}

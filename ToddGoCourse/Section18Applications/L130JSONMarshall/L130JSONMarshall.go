package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   27,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(string(bs))
	os.Stdout.Write(bs)
}

//func Marshal(v interface{}) ([]byte, error)

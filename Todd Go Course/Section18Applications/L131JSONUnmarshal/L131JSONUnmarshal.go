package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":27},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s) // conversion

	var people []person

	fmt.Println(s, bs)

	fmt.Printf("\n%T", bs)
	fmt.Printf("\n%T", s)

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("\n all data", people)

	for i, v := range people {
		fmt.Println("Position", i, v)
	}

}

//We can take JSON and bring it back into our Go program by unmarshalling that JSON.
//func Unmarshal(data []byte, v interface{}) error

//func Marshal(v interface{}) ([]byte, error)

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	bs := []byte(s)

	fmt.Println(bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)
	var people []person

	err := json.Unmarshal(bs, &people) // could have not dome bs conv above and just done err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("all data", people)

	for i, v := range people {
		fmt.Println("Person no:", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for _, v2 := range v.Sayings {
			fmt.Println("\t\t", v2)
		}
	}

	//func Unmarshal(data []byte, v interface{}) error

}

//Hands-on exercise #2
//Starting with this code, unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/

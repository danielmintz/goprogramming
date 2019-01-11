package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	Favflav []string
}

func main() {

	//m1 := map[string][]string{
	//	`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
	//	`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
	//	`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	//}

	p1 := person{
		first: "James",
		last:  "Bond",
		Favflav: []string{
			"choco",
			"strawberry",
			"rum and raisin",
		},
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		Favflav: []string{
			"vanilla",
			"coffee choc",
			"hazelnut",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {

		fmt.Println(v.first)
		fmt.Println(v.last)
		for i2, v2 := range v.Favflav {
			fmt.Println(i2, v2)
		}
		fmt.Println("--------")
	}
}

//Take the code from the previous exercise,
//then store the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.

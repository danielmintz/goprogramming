package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`mintz_daniel`] = []string{`alpro, unsweetened`, `MTB`, `Learning`}

	delete(m, `no_dr`)

	//for i, v := range m {
	//	fmt.Println(i, v)
	//}
	//fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

//Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

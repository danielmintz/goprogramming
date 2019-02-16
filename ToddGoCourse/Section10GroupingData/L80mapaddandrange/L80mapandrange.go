package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 42,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Moneypenny"])
	fmt.Println(m["Barnabas"])
	m["Todd"] = 33
	fmt.Println(m["Todd"])

	v, ok := (m["James"])
	fmt.Println(v, ok)

	if v, ok := m["James"]; ok {
		fmt.Println("If comma OK is corect on map query then print", v)

	}
	for i, v := range m {
		fmt.Println(i, v)

	}
}

//Map - add element & range
//Here is how we add an element to a map. I also show you how to use the range loop to print out a map’s keys and values.

package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 42,
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Moneypenny"])

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("this is the case", v)
		delete(m, "Moneypenny")
		fmt.Println(m)
	}
}

//Map - delete
//You delete an entry from a map using delete(<map name>, “key”). No error is thrown if you use
//a key which does not exist.
//To confirm you delete a key/value, verify that the key/value exists with the comma ok idiom.

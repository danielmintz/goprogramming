package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      29,
		"Moneypenny": 24,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Moneypenny"])
	fmt.Println(m["barnabas"])

	//v, ok := m["James"]
	//fmt.Println(v)
	//fmt.Println(ok)

	//v, ok := m["Barnabas"]
	//fmt.Println(v)
	//fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("this is the if print", v)
	}
}

//A map is a key-value store. This means that you store some value and you access that value by a key.
//For instance, I might store the value “phoneNumber” and access it be the key “friendName”.
//This way I could enter my friend’s name and have the map return their phone number.
//The syntax for a map is map[key]value. The key can be of any type which allows comparison
//(eg, I could use a string or an int, for example, as I can compare if two strings are equal,
//or if two ints are equal). It is important to note that maps are unordered.
//If you print out all of the keys and values in a map, they will print out in random order.
//The comma ok idiom is also covered in this video. A map is the perfect data structure when you
//need to look up data fast.

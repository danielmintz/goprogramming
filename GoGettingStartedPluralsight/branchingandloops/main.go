package main

import "fmt"

func main() {
	// branch
	foo := 2
	if foo == 1 {
		fmt.Println("you rock")
	} else {
		fmt.Println("you suck")
	}

	if bar := 2; bar == 2 {
		fmt.Println("again you rock")
	}

	switch foo := 2; foo {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")

	}
	foobar := 7
	switch {
	case foobar == 1:
		fmt.Println("One")
	case foobar > 1:
		fmt.Println("yup it's more than foo")

	}

	// init condition and loop

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	x := 0
	for {
		x++
		fmt.Println(x)
		if x > 5 {
			break
		}
	}
	// Ranging over a slice of an array
	s := []string{"hey", "hello", "Shalom", "Yo"}
	for i, v := range s {
		fmt.Println(i, v)
	}

	//maps and ranging over the key and the entry

	m := make(map[string]string)
	m["first"] = "James"
	m["second"] = "Bond"
	m["third"] = "London"
	for k, v := range m {
		fmt.Println(k, v)
	}
}

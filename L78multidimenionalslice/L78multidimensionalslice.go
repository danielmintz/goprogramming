package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Peach"}
	fmt.Println(mp)

	mp = append(mp, "avocado")
	fmt.Println(mp)

	q := [][]string{jb, mp}
	fmt.Println(q)
}

//A multi-dimensional slice is like a spreadsheet. You can have a slice of a slice of some type.

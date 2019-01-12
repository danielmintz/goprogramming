package main

import (
	"fmt"
)

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xxs := [][]string{x, y}

	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t %v \n", j, val)
		}
	}

}

//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."

//Range over the records, then range over the data in each record.

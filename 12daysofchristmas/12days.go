package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Println("On the", day, "of Christmas my true love sent to me:")
		if day == 12 {
			fmt.Println("Twelve Drummers drumming")
		} else if day == 11 {
			fmt.Println("Elven Pipers Piping")
		} else if day == 10 {
			fmt.Println("Ten Lords a Leaping")
		} else if day == 9 {
			fmt.Println("Nine ladies dancing")
		} else if day == 8 {
			fmt.Println("Eight maids a-milking")
		} else if day == 7 {
			fmt.Println("Seven swans a-swimming")
		} else if day == 6 {
			fmt.Println("Six geese a-laying")
		} else if day == 5 {
			fmt.Println("Five golden rings")
		} else if day == 4 {
			fmt.Println("Four calling birds")
		} else if day == 3 {
			fmt.Println("Three French hens")
		} else if day == 2 {
			fmt.Println("Two turtle doves")
		} else if day == 1 {
			fmt.Println("And a partridge in a pear tree!")
		}

	}

}

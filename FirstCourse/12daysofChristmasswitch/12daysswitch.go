package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Println("On the", day, "of Christmas my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("Twelve Drummers drumming")
		case 11:
			fmt.Println("Elven Pipers Piping")
		case 10:
			fmt.Println("Ten Lords a Leaping")
		case 9:
			fmt.Println("Nine ladies dancing")
		case 8:
			fmt.Println("Eight maids a-milking")
		case 7:
			fmt.Println("Seven swans a-swimming")
		case 6:
			fmt.Println("Six geese a-laying")
		case 5:
			fmt.Println("Five golden rings")
		case 4:
			fmt.Println("Four calling birds")
		case 3:
			fmt.Println("Three French hens")
		case 2:
			fmt.Println("Two turtle doves")
		case 1:
			fmt.Println("And a partridge in a pear tree!")

		}

	}
}

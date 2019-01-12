package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("")
	fmt.Println("On the first day of Christmas my true love sent to me")
	fmt.Println("a Partridge in a Pear Tree")
	fmt.Println("")
	time.Sleep(time.Second)

	for day := 2; day <= 12; day++ {
		fmt.Println("On the", day, "day of Christmas my true love sent to me:")
		switch day {
		case 12:
			fmt.Println("Twelve Drummers drumming")
			time.Sleep(time.Second)
			fallthrough
		case 11:
			fmt.Println("Elven Pipers Piping")
			time.Sleep(time.Second)
			fallthrough
		case 10:
			fmt.Println("Ten Lords a Leaping")
			time.Sleep(time.Second)
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing")
			time.Sleep(time.Second)
			fallthrough
		case 8:
			fmt.Println("Eight maids a-milking")
			time.Sleep(time.Second)
			fallthrough
		case 7:
			fmt.Println("Seven swans a-swimming")
			time.Sleep(time.Second)
			fallthrough
		case 6:
			fmt.Println("Six geese a-laying")
			time.Sleep(time.Second)
			fallthrough
		case 5:
			fmt.Println("Five golden rings")
			time.Sleep(time.Second)
			fallthrough
		case 4:
			fmt.Println("Four Calling Birds")
			time.Sleep(time.Second)
			fallthrough
		case 3:
			fmt.Println("Three French hens")
			time.Sleep(time.Second)
			fallthrough
		case 2:
			fmt.Println("Two turtle doves")
			time.Sleep(time.Second)
			fallthrough
		case 1:
			fmt.Println("And a partridge in a pear tree!")
			time.Sleep(time.Second)

		}
		fmt.Println("")

	}
}

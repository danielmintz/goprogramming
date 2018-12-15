package main

import "fmt"

const gCap int = 5          // Total capacity of our grocery list
var gGroceries [gCap]string //think string means words so can be up to 5 groceries
var gLen int                //Total number of grocery items in our current array

func addGrocery(a string) {
	if gLen < gCap {
		gGroceries[gLen] = a
		gLen++ // addGrocery means you can keep adding a grocery item one by one up to 5 or else the below error....
	} else {
		fmt.Println("too many items, we are done for!!")
	}

}

func listGroceries() {
	fmt.Println("Grocery list as follows:")
	for i := 0; i < gLen; i++ { //don't get whay this isn't gCap? but basically saying PRINT each of the list one by one.
		fmt.Println(gGroceries[i]) //
	}

}
func main() {

	addGrocery("Bread")
	addGrocery("Salt")
	addGrocery("Vinegar")
	addGrocery("Butter")
	addGrocery("Milk")
	addGrocery("ice cream")
	addGrocery("Ketchup")

	listGroceries() //run the listgrocery function below which prints it out and runs the functions.

}

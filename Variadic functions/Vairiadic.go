package main

import "fmt"

var gGroceries []string // string is like words

func addGrocery(a ...string) { //three dots makes it a Variadic function (i.e. means a lot of.. more than 2 i.e a range
	fmt.Println("Capacity", cap(gGroceries)) //indicates capacity of our slice i.e. just showing how it work (can comment out)
	for _, d := range a {                    //This is the forloop for d  the data forfor every element of the slice of strings passed in a

		gGroceries = append(gGroceries, d) //this is the slice command bit which means keep growing the array as bits

	} // are added Magic apend value newslice:= append(oldslice,newvalue)

}

func listGroceries() {
	fmt.Println("Grocery list as follows:")
	/*for i := 0; i < len(gGroceries); i++ { //d saying PRINT each of the list one by one.
				fmt.Println(gGroceries[i]) //
	}*/
	for _, d := range gGroceries { // This is the FOR LOOP doing what line commente dout does
		// but more efficiently .... i.e. print each one of the list one by one.
		fmt.Println(d)
	}
}
func main() {

	addGrocery("Bread", "Salt", "Vinegar", "Butter", "Milk", "ice cream", "Ketchup", "Balsamic Vinegar")

	listGroceries() //run the listgrocery function below which prints it out and runs the functions.

}

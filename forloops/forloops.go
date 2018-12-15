package main

import "fmt"

var gGroceries []string // string is like words

func addGrocery(a string) {
	fmt.Println("Capacity", cap(gGroceries)) //indicates capacity of our slice i.e. just showing how it work (can comment out)
	gGroceries = append(gGroceries, a)       //this is the slice command bit which means keep growing the array as bits
	// are added Magic apend value newslice:= append(oldslice,newvalue)

}

func listGroceries() {
	fmt.Println("Grocery list as follows:")
	/*for i := 0; i < len(gGroceries); i++ { //d saying PRINT each of the list one by one.
				fmt.Println(gGroceries[i]) //
	}*/
	for _, d := range gGroceries { // This is the FOR LOOP doing what line above commented out does
		// but more efficiently
		fmt.Println(d)
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
	addGrocery("Balsamic Vinegar")
	addGrocery("Balsamic Vinegar1")
	addGrocery("Balsamic Vinegar2")
	addGrocery("Balsamic Vinegar3")
	addGrocery("Balsamic Vinegar4")
	addGrocery("Balsamic Vinegar5")

	listGroceries() //run the listgrocery function below which prints it out and runs the functions.

}

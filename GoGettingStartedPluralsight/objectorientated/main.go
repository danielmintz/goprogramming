package main

import "fmt"

type myStruct struct {
	myField string
}

// fifth way uses an ampersand '&' but avoids using new fucntion

func main() {
	foo := &myStruct{"bar"}
	fmt.Println(foo.myField)
}

// fourth way uses the "new" func and uses poitners for multiple returns

// func main() {
// 	foo := new(myStruct)
//	foo.myField = "bar"
//	fmt.Println(foo)
// }

//third way just a tiny bit cleaner than second way

//func main() {

//	foo := myStruct{"bar"}
//	fmt.Println(foo.myField)
// }

//second way to create an object--

// func main() {
//	foo := myStruct{
//		myField: "Bar",
//	}
//	fmt.Println(foo.myField)
//}

// first way to create an object--

// func main() {
//	foo := myStruct{}
//	foo.myField = "bar"
//	fmt.Println(foo.myField)

//}

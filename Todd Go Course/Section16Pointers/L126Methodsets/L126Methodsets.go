package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {

	c := circle{
		radius: 5,
	}

	//info(c)  //NON-POINTER RECEIVER & NON-POINTER VALUE
	//info(&c) // NON-POINTER RECEIVER & POINTER VALUE

	fmt.Println(c.area()) // method set determines the INTERFACES that the type implements

}

//Method sets
//Method sets determine what methods attach to a TYPE.
//It is exactly what the name says: What is the set of methods for a given type?
//That is its method set.

//IMPORTANT: “The method set of a type determines the INTERFACES that the type implements.....”

//a NON-POINTER RECEIVER
//works with values that are POINTERS or NON-POINTERS.
//a POINTER RECEIVER
//only works with values that are POINTERS.

//Receivers       Values
//-----------------------------------------------
//(t T)           T and *T
//(t *T)          *T

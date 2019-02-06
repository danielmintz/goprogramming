package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByName implements sort.Interface for []Person based on
// the Name field.
type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {
	p1 := Person{"James", 27}
	p2 := Person{"Dr No", 56}
	p3 := Person{"Miss Moneypenny", 26}
	p4 := Person{"Q", 54}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(ByAge(people)) // this uses conversion to convert people to type by age.
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}

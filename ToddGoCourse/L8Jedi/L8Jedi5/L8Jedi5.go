package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge implements sort.Interface for []user based on
// the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast implements sort.Interface for []user based on
// the Last field.
type ByLast []user

func (bs ByLast) Len() int           { return len(bs) }
func (bs ByLast) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs ByLast) Less(i, j int) bool { return bs[i].Last < bs[j].Last }

// BySayings implements sort.Interface for []user based on
// the Sayings field.
//type BySayings []user

//func (xi BySayings) Len() int           { return len(xi) }
//func (xi BySayings) Swap(i, j int)      { xi[i], xi[j] = xi[j], xi[i] }
//func (xi BySayings) Less(i, j int) bool { return xi[i].Sayings < xi[j].Sayings }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	//	fmt.Println("\n", users)

	// your code goes here

	sort.Sort(ByAge(users))
	//fmt.Println("\n", users)
	for i, v := range users {
		fmt.Println("Person number:", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, vs := range v.Sayings {
			fmt.Println("\t\t", vs)
		}
	}

	sort.Sort(ByLast(users))
	//fmt.Println("\n", users)
	for i, v := range users {
		fmt.Println("Person number:", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		//sort.Strings(v.Sayings)
		for _, vs := range v.Sayings {
			fmt.Println("\t\t", vs)
		}

	}

	//sort.Sort(ByLast(users))
	// fmt.Println("\n", users)
	//sort.Strings(xs)

}

//Starting with this code, sort the []user by
//age
//last
//Also sort each []string “Sayings” for each user
//print everything in a way that is pleasant

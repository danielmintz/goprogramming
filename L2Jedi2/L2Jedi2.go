package main

import "fmt"

func main() {

	g := (50 == 30)
	fmt.Println(g)
	h := (50 <= 30)
	fmt.Println(h)
	i := (50 >= 30)
	fmt.Println(i)
	j := (50 != 30)
	fmt.Println(j)
	k := (50 > 30)
	fmt.Println(k)
	l := (50 < 30)
	fmt.Println(l)
	fmt.Printf("%T\n%T\n", g, h)

}

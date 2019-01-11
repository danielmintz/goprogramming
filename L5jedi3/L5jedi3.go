package main

import (
	"fmt"
)

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	p1 := truck{
		fourWheel: true,
		vehicle: vehicle{
			doors:  6,
			colour: "red",
		},
	}

	p2 := sedan{
		luxury: false,
		vehicle: vehicle{
			doors:  4,
			colour: "silver",
		},
	}
	fmt.Println(p1.fourWheel, p1.doors, p1.colour)
	fmt.Println(p2.luxury, p2.doors, p2.colour)

}

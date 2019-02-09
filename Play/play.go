package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	Loginword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(Loginword1))
	if err != nil {
		fmt.Println("You can not log in incorrect password")
		return
	}
	fmt.Println("you are loged in")
}

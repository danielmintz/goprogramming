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

	LoginPword1 := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(LoginPword1))
	if err != nil {
		fmt.Println("you cant log in")
		return
	}
	fmt.Println("you are loggen in")

}

//bcrypt

//func GenerateFromPassword(password []byte, cost int) ([]byte, error)

//All too often today you still hear about websites and databases being hacked and user’s information being compromised.
//There is no excuse for this. As programmers, we have the tools to protect user data. Bcrypt is
//one of the tools you can use to protect user data.

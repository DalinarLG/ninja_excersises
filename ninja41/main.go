package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `123456`
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bytes))

	login := `123456`
	err = bcrypt.CompareHashPassword(bytes, []byte(login))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("login success")

}

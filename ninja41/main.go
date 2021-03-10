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

	login := `1234568`
	err = bcrypt.CompareHashAndPassword(bytes, []byte(login))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("login success")

}

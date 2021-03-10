package main

import (
	"encoding/json"
	"fmt")


type user struct {
	Name string
	Age  int
}

func main() {
	us := user{
		Name: "Leonardo",
		Age:  37,
	}

	us2 := user{
		Name: "Jared",
		Age:  6,
	}

	users := []user{us, us2}	
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
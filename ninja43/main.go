package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	Name string //`json:"name"`
	Age int		//`json:"int"`
}

func main(){
	s := `[{"Name":"Leonardo","Age":37},{"Name":"Jared","Age":6}]`

	var users []user	
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(users)
}
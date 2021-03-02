package main

import "fmt"

type person struct {
	firstname string
	lastname string
	favoriteFlavor []string
}
func main() {

	p1 := person{
		firstname: "Leonardo",
		lastname: "Guerrero",
		favoriteFlavor:[]string{
			"vanilla",
			"chocolate",
			"strawberry",
		},
	}
	fmt.Println(p1.firstname, p1.lastname)
	for _, flavor := range p1.favoriteFlavor{
		fmt.Printf("\t %s\n", flavor)
	}
}

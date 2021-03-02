package main

import "fmt"

type person struct {
	firstname string
	lastname string
	favoriteFlavor []string
}

var people = map[string]person{}

func main() {
	people["guerrero"] = person{
		firstname: "Leonardo",
		lastname: "Guerrero",
		favoriteFlavor:[]string{
			"vanilla",
			"chocolate",
			"strawberry",
		},

	}

	people["laura"] = person{
		firstname: "Laura",
		lastname: "Gonzalez",
		favoriteFlavor:[]string{
			"shit",
			"poo",
			"dookie",
		},
	}

	for _, p := range people{
		fmt.Printf("%s %s \n",p.firstname, p.lastname)
		for j, f := range p.favoriteFlavor{
			fmt.Printf("\t%v %s\n",j, f)
		}
	}
}
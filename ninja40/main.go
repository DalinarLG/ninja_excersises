package main

import "fmt"

type person struct {
	name     string
	lastname string
	age      int
}

func changeMe(p *person) {
	p.name = "Brandon"
	p.lastname = "Sanderson"
	p.age = 45
}
func main() {

	person := person{
		name:     "Leonardo",
		lastname: "Guerrero",
		age:      37,
	}

	fmt.Println(person)
	changeMe(&person)
	fmt.Println(person)
}

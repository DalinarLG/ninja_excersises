package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func (p person) speak() {
	fmt.Println("My name is ",p.First," and I'm ",p.Age," years old")
}

func main(){
	p := person{
		First: "Leonardo",
		Last: "Guerrero",
		Age: 37,
	}

	p.speak()
}
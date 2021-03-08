package main

import "fmt"



func main() {

	name := func()string {
		return "Hello my name is Leonardo"
	}
	
	fmt.Println(name())
	fmt.Printf("%T",name)
}

package main

import "fmt"

type powermetal int

var x powermetal

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

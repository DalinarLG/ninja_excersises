package main

import "fmt"



func main() {

	arr := [5]int{4,5,6,2,3}
	for i, v := range arr{
		fmt.Printf("%v -- %v\n", i , v)
	}

	fmt.Printf("%T", arr)
}

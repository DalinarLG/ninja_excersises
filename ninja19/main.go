package main

import "fmt"

func main() {
	xi := []int{78, 789, 21, 45, 2, 4, 5, 6, 9, 10}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", xi)
}

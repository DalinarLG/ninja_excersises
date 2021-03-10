package main

import (
	"fmt"
	"sort")


func main() {
	xi := []int{4, 5, 2, 148, 1213, 41, 2, 14, 78, 663, 2}
	xs := []string{"leonard", "jared", "ashley", "lauara", "hanny"}

	sort.Ints(xi)
	fmt.Println(xi)

	sort.Strings(xs)
	fmt.Println(xs)
}
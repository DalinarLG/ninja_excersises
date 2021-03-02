package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xi := append(x[:5])
	xii := append(x[5:])
	xiii := append(x[2:7])
	xiiii := append(x[1:6])

	fmt.Println(xi)
	fmt.Println(xii)
	fmt.Println(xiii)
	fmt.Println(xiiii)
}

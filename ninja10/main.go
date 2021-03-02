package main

import "fmt"

var x = 7

func main() {

	fmt.Printf("%d\t%b\t%x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%x", y, y, y)

}

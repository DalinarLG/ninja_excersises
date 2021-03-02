package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		x := i / 4
		y := i % 4
		fmt.Printf("%d Divided by 4: %d  Modulus %d\n", i, x, y)
	}
}
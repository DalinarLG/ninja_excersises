package main

import "fmt"

func main() {
	x := 1983
	for {
		if x <= 2021 {
			fmt.Println(x)
		}
		x++
	}
}
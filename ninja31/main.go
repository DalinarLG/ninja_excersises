package main

import "fmt"

func foo (xi ...int)int{
	sum := 0
	for _, v := range xi{
		sum += v
	}
	return sum
}

func bar (x []int)int{
	sum := 0
	for _, v := range x{
		sum += v
	}
	return sum
}

func main() {
	xi := []int{1,5,6,7,8,9,2}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}
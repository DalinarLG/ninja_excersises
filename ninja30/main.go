package main

import "fmt"

func foo() int {
	return 37
}
func bar()(int, string) {
	return 42, "hello"
}
func main() {
	x := foo()
	y, z := bar()
	fmt.Println(x,y,z)

}
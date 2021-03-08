package main

import "fmt"

func foo() {
	fmt.Println("func foo")
}

func bar(){
	fmt.Println("Func bar")
}

func main() {

	defer foo()
	bar()
}
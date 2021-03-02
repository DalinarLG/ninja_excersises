package main

import (
	"fmt"
	"runtime"
)


const (
	a = iota
	b = iota
	c = iota
	d = iota
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
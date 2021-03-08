package main

import (
	"fmt"
	"math"
)

type circle struct {
	r float64
}

type square struct {
	height float64
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (s square) Area() float64 {
	return math.Pow(s.height, 2)
}

type shapes interface {
	Area()float64
}

func info(sh shapes) {
	fmt.Println(sh.Area())
}

func main() {

	c := circle{
		r: 12.345,
	}

	s := square{
		
		height: 15,
	}

	info(c)
	info(s)

}

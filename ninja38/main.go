package main


import "fmt"

func foo(f func(xi [] int)int, xii []int)int{
	n := f(xii)
	n++
	return n
}

func main(){
	g := func(xi []int)int{
		if len(xi) == 0{
			return 0
		}
		if len(xi) == 1{
			return xi[0]
		}

		return xi[0]+ xi[len(xi)-1]
	}
	xi := []int{1,2,3,4,5,6}
	fmt.Println((g(xi)))
}
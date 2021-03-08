package main


import "fmt"

func bar()func()string{
	return func() string {
		return "hello theree"
	}
}

func main(){

	v:= bar()
	fmt.Println(v())

}
package main

import "fmt"

func main(){

	j := []string{
		"James",
		"Bond",
		"Shaken, not stirred",
	}

	m := []string{
		"Miss",
		"MoneyPenny",
		"Hellooooo, James",
	}

	xi := [][]string{j, m}

	fmt.Println(xi)
	for i, v := range xi{
		fmt.Printf("Tupla %v\n", i)
		for j, xs := range v{
			fmt.Printf("\t%v - %s\n",j,xs)
		}
	}
}
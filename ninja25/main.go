package main

import "fmt"

func main() {
	
	people := map[string][]string{}

	people["bond_james"] = []string{"Shaken, not stirred", "martinis", "women"}
	people["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	people["no_dr"] = []string{"Being evil", "Sunsets", "Ice cream"}

	fmt.Println(people)

	for k, value := range people{
		fmt.Printf("Character: %s \n", k)
		for j, val := range value{
			fmt.Printf("\t %v - %s \n ", j, val)
		}
		

	}
}
package main

import "fmt"

func main() {
	people := map[string][]string{}

	people["bond_james"] = []string{"Shaken, not stirred", "martinis", "women"}
	people["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	people["no_dr"] = []string{"Being evil", "Sunsets", "Ice cream"}

	delete(people, "no_dr")

	fmt.Println(people)
}
package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name     string
	Lastname string
	Age      int
	Fruits   []string
}

//ByAge sort
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//ByLastname sort
type ByLastname []user

func (l ByLastname) Len() int           { return len(l) }
func (l ByLastname) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByLastname) Less(i, j int) bool { return l[i].Lastname < l[j].Lastname }

func main() {
	us := user{
		Name:     "Leonardo",
		Lastname: "Guerrero",
		Age:      37,
		Fruits:  []string{"Mango", "Melon", "Patilla"},
	}

	us2 := user{
		Name:     "Jared",
		Lastname: "Lara",
		Age:      6,
		Fruits:  []string{"Manzana", "Naranja", "Fresa"},
	}

	us3 := user{
		Name:     "Catalina",
		Lastname: "Gonzalez",
		Age:      54,
		Fruits:  []string{"Maracuya", "Tomate", "Papaya"},
	}

	us4 := user{
		Name:     "Kelly",
		Lastname: "Prada",
		Age:      40,
		Fruits:  []string{"verga", "mondá", "picha"},
	}

	users := []user{us, us2, us3, us4}

	sort.Sort(ByAge(users))
	for _, u := range users {
		fmt.Println(u.Name, u.Lastname, u.Age )
		sort.Strings(u.Fruits)
		for _, s := range u.Fruits{
			fmt.Printf("\t %v\n", s)
		}
	}

	fmt.Println("-----------------------")
	sort.Sort(ByLastname(users))
	for _, u := range users {
		fmt.Println(u.Name, u.Lastname, u.Age )
		for _, s := range u.Fruits{
			fmt.Printf("\t %v\n", s)
		}
	}

}

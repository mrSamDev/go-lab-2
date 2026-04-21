package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favIcecreams []string
}

func main() {
	store := map[string]person{}
	p1 := person{
		firstName:    "Alice",
		lastName:     "Johnson",
		favIcecreams: []string{"Vanilla", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

	store[p1.lastName] = p1
	fmt.Printf("Stored person: %v\n", store[p1.lastName])

}

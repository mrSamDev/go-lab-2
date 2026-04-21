package main

import "fmt"

type person struct {
	firstName    string
	lastName     string
	favIcecreams []string
}

func main() {

	p1 := person{
		firstName:    "Alice",
		lastName:     "Johnson",
		favIcecreams: []string{"Vanilla", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)

}

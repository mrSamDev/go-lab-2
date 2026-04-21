package main

import "fmt"

func main() {

	p1 := struct {
		firstName string
		friends   map[string]string
		favDrinks []string
	}{
		firstName: "Alice",
		friends: map[string]string{
			"Bob":   "Best friend",
			"Carol": "Close friend",
		},
		favDrinks: []string{"Coffee", "Tea", "Milk"},
	}

	fmt.Printf("%T\n", p1)
}

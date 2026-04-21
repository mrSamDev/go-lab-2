package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type secretPerson struct {
	Person
	Secret string
}

func main() {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Bob", Age: 25}
	s1 := secretPerson{
		Person: Person{Name: "Charlie", Age: 35},
		Secret: "Likes pineapple on pizza",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(s1)
	fmt.Printf("%T\n", p1)

}

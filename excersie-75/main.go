package main

import "fmt"

type english struct {
	greeting string
}

func main() {

	a := english{greeting: "Hello, World!"}

	fmt.Println(a.greeting)

	fmt.Println(changeName(a, "Hello, Universe!").greeting)

	fmt.Println(a.greeting)
}

func changeName(e english, newName string) english {
	e.greeting = newName
	return e
}

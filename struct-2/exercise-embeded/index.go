package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine engine
	make   string
	model  string
	doors  int
	color  string
}

func main() {
	v1 := vehicle{
		engine: engine{electric: true},
		make:   "Tesla",
		model:  "Model S",
		doors:  4,
		color:  "Red",
	}

	v2 := vehicle{
		engine: engine{},
		make:   "Ford",
		model:  "Mustang",
		doors:  2,
		color:  "Blue",
	}
	fmt.Println(v1)
	fmt.Println(v2)
}

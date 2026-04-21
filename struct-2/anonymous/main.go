package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{name: "Alice", age: 30}
	fmt.Printf("%T\n", p)

}

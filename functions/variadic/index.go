package main

import (
	"fmt"
)

func printNames(firstname string, names ...string) {
	fmt.Println("First name:", firstname)
	for _, name := range names {
		fmt.Println(name)
	}
}

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	printNames("Alice", "Bob", "Charlie")
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
}

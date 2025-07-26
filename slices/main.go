package main

import (
	"fmt"

	icecream "example.com/price-calculator/go-learning/go-path-2/slices/ice-cream"
)

func main() {

	//array
	arr := [...]string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("Original array:", arr)
	fmt.Println("Length of array:", len(arr))
	fmt.Printf("%T\n", arr)

	//slices
	icecream.IceCream()

}

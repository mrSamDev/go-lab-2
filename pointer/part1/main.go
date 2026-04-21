package main

import "fmt"

func main() {

	var a int = 10

	b := &a

	*b = 12

	fmt.Println(a)

	fmt.Println(b)

	fmt.Printf("type %T\n", *b)

	fmt.Println(&a)

}

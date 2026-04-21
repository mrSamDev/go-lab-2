package main

import "fmt"

type english struct {
	greeting string
}

var a, b, c *string

type human struct {
	rage string
}

type killWithKindness interface {
}

type killWithRage interface {
}

func main() {

	q := "we"
	w := "are"
	e := "learning"

	a = &q
	b = &w
	c = &e

	x := []*string{a, b, c}

	for _, n := range x {
		fmt.Println("%v \t %T", n, *n)
	}

	sijo := english{greeting: "Hello, World!"}
	sam := "SEWMOP"

	pointerToSijo := &sijo

	// print the memory address
	fmt.Println(&sam)
	fmt.Printf("%p\n", pointerToSijo)

}

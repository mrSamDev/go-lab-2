package main

import "fmt"

type Color map[string]string

func main() {

	// var xColor Color

	// mapColor := make(Color)

	color := Color{"red": "#FF0000", "blue": "#0000FF", "white": "#FFFFFF"}
	color.print()

	// xColor = make(Color)

	// xColor["red"] = "dsad"

	// mapColor["red"] = "Dsada"

	// delete(xColor, "red")

	// fmt.Println(color, xColor, mapColor)
}

func (c Color) print() {
	for key, value := range c {
		fmt.Printf("key: %v ----- value: %v \n", key, value)
	}
}

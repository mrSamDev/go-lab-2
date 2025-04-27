package main

import "fmt"

const (
	Red = iota // Red is 0
	Blue
	Green
)

type BitSize int

const (
	_          = iota
	KG BitSize = 1 << (10 * iota)
	MB BitSize = 1 << (10 * iota)
	GB BitSize = 1 << (10 * iota)
)

func main() {
	var highUint8 uint8 = 255
	var highInt8 int8 = 127

	fmt.Printf("%b \t \t \b %b \n", KG, KG, highInt8)
	fmt.Printf("%b \t \t \b %b", MB, MB, highUint8)
}

package xshape

import "fmt"

type shape interface {
	area() float64
}

type triange struct {
	height float64
	base   float64
}

func (t triange) area() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	sideLength float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	newTriangle := triange{
		height: 10.00,
		base:   2.00,
	}

	newSquare := square{
		sideLength: 2.00,
	}

	printArea(newTriangle)
	printArea(newSquare)

}

func printArea(s shape) {
	area := s.area()
	fmt.Println("[area] is eqaul to", area)
}

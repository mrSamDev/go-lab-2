package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	fmt.Println(split(19))
}

// 0 > 0
// 1 > 1
// 2 > 01
// 3 > 11
// 4 > 100
// 5 > 101
// 6 > 110
//7 > 111
//8 > 1000
//9 > 1001
//
//11 > 1011
//12 > 1100

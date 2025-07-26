package main

import "fmt"

func main() {
	arr := [10]int{}

	// min := 42
	// i := 0

	// for {
	// 	arr[i] = min
	// 	i++
	// 	min++

	// 	if min > 51 {
	// 		break
	// 	}

	// }

	for i := 0; i < len(arr); i++ {
		arr[i] = 42 + i
	}
	for i, v := range arr {
		fmt.Println(i, v) // Output: 0 42, 1 43, ..., 9 51
	}

	fmt.Println(len(arr)) // Output: 10
	fmt.Println(cap(arr)) // Output: 10
	fmt.Println(arr)      // Output: [42 43 44 45 46 47 48 49 50 51]
}

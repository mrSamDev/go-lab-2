package main

import "fmt"

// main: Function to demonstrate basic operations on an array/slice in go.
func main() {

	arr := make([]int, 5)

	fmt.Println("\nLength of array:", len(arr))
	fmt.Println("Capacity of array:", cap(arr))

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1 // Initializing the array with values 1 to 5
	}

	fmt.Println("Array elements:", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %T \n", arr[i], arr[i])
	}
	fmt.Println("\nLength of array:", len(arr))
	fmt.Println("Capacity of array:", cap(arr))

}

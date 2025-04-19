package main

import "fmt"

func main() {

	numbers := generateNumber(100)

	fmt.Println(numbers)

	for _, number := range numbers {

		if number%2 == 0 {
			fmt.Printf("%v the number is even \n", number)
		} else {
			fmt.Printf("%v the number is odd \n", number)
		}

	}

}

func generateNumber(limit int) []int {

	var numbers = make([]int, limit+1)

	for i := 0; i <= limit; i++ {
		numbers[i] = i
	}

	return numbers

}

func generateNumberWithoutMake(limit int) []int {

	var numbers = []int{}

	for i := 0; i <= limit; i++ {
		numbers = append(numbers, i)
	}

	return numbers

}

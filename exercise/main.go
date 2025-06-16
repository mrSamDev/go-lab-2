package main

import (
	"fmt"
	"math/rand"
)

func checker() {
	randomvalue := rand.Intn(250)
	fmt.Printf("Random value is: %d\n", randomvalue)
	switch {
	case randomvalue > 0 && randomvalue <= 50:
		fmt.Println("The random value is between 1 and 50")
	case randomvalue > 50 && randomvalue <= 100:
		fmt.Println("The random value is between 51 and 100")
	case randomvalue > 100 && randomvalue <= 150:
		fmt.Println("The random value is between 101 and 150")
	case randomvalue > 150 && randomvalue <= 200:
		fmt.Println("The random value is between 151 and 200")
	case randomvalue > 200:
		fmt.Println("The random value is gtreater than 200")
	}
}

func main() {
	z := 0

	for {
		checker()
		z++
		if z >= 100 {
			fmt.Println("Exiting the loop after 10 iterations.")
			break
		}
	}
}

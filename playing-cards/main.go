package main

import (
	"fmt"

	"example.com/playing-cards/decks"
)

func main() {
	cards := decks.NewDeck()
	fmt.Println(len(cards))
	cards.Shuffle()
	// x, _ := decks.Deal(cards, 4)
	cards.Display()
	// x.Save("x.txt")
	// data, error := x.Read("x.txt")

	// if error != nil {
	// 	fmt.Println(error)
	// 	os.Exit(1)
	// }
	// data.Display()

	// data.Display()
	// deal2.Display()

}

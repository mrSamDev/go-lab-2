package decks

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"example.com/playing-cards/utils"
)

type Deck []string

func NewDeck() Deck {
	suits := map[string]string{
		"Spades":   "♠",
		"Hearts":   "♥",
		"Diamonds": "♦",
		"Clubs":    "♣",
	}

	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	var deck Deck
	for _, symbol := range suits {
		for _, value := range values {
			deck = append(deck, fmt.Sprintf("%s%s", value, symbol))
		}
	}
	return deck
}

func (d Deck) Display() {

	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d Deck) ToString() string {

	return strings.Join(d, ",")

}

func (d Deck) Shuffle() Deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range d {
		newPosition := r.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

func (d Deck) Save(fileName string) error {
	return utils.SaveFile(fileName, utils.ConvertToByteSize(d.ToString()))
}

func (d Deck) Read(fileName string) (Deck, error) {
	data, error := utils.ReadFile(fileName)
	if error != nil {
		return nil, error
	}

	splited_value := strings.Split(data, ",")

	return Deck(splited_value), nil
}

func Deal(d Deck, handsize int) (Deck, Deck) {
	return d[:handsize], d[handsize:]
}

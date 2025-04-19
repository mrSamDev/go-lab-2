package decks_test

import (
	"testing"

	"example.com/playing-cards/decks"
	"example.com/playing-cards/utils"
)

func TestNewDeck(t *testing.T) {
	decks := decks.NewDeck()

	if len(decks) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(decks))
	}

	if decks[0] != "A♠" {
		t.Errorf("Expected 1st card to be A♠, but got %v", decks[0])
	}

	if decks[len(decks)-1] != "K♣" {
		t.Errorf("Expected last card to be K♣, but got %v", decks[len(decks)-1])
	}

}

var testing_filename = "_decktesting.text"

func TestSaveFileAndReadTheFile(t *testing.T) {
	utils.RemoveFile(testing_filename)
	decks := decks.NewDeck()
	error := decks.Save(testing_filename)

	if error != nil {
		t.Errorf("Saving to file failed %v", testing_filename)
	}

	load_deck, error := decks.Read(testing_filename)

	if error != nil {
		t.Errorf("Umable to open the saved file %v", testing_filename)
	}

	if len(load_deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(load_deck))
	}

	utils.RemoveFile(testing_filename)
}

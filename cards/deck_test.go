package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLength := 16

	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedLength, len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	expectedLength := 16

	if len(loadedDeck) != expectedLength {
		t.Errorf("Expected %v cards in deck, got %v", expectedLength, len(loadedDeck))
	}

	os.Remove("_decktesting")
}

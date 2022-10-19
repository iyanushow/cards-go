package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFIle(t *testing.T) {
	os.Remove("_decktest")

	deck := newDeck()

	deck.saveToFile("_decktest")

	loadedDeck := newDeckFromFile("_decktest")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck of length 52, but got %v", len(deck))
	}

	os.Remove("_decktest")
}

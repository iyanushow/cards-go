package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of deck which is a string slice
type deck []string

func (d deck) printCards() {
	for _, card := range d {
		fmt.Print(card)
	}
}

func (d deck) toString() string {
	deckSlice := []string(d)

	return strings.Join(deckSlice, ",")
}

func (d deck) saveToFile(filename string) error {
	deckByteSlice := []byte(d.toString())

	return os.WriteFile(filename, deckByteSlice, 0666)
}

func newDeck() deck {
	cards := deck{}

	cardSuits := [4]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := [13]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen", "Joker"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	bs_string := strings.Split(string(bs), ",")

	return deck(bs_string)
}

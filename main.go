package main

func main() {

	// cards := newDeck()

	// // hand, remainingDeck := deal(cards, 5)

	// // hand.printCards()
	// // remainingDeck.printCards()

	// cards.saveToFile("newCards")

	cards := newDeckFromFile("newCards")
	cards.shuffle()
	cards.printCards()

}

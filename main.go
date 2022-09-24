package main

func main() {

	// cards := newDeck()

	// // hand, remainingDeck := deal(cards, 5)

	// // hand.printCards()
	// // remainingDeck.printCards()

	// cards.saveToFile("newCards")

	cards := newDeckFromFile("newCardss")

	cards.printCards()

}

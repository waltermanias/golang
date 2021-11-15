package main

func main() {
	// deck := newDeck()
	// // deck.print()

	// hand, remainingDeck := deck.deal(5)

	// hand.print()
	// fmt.Println("--------")
	// remainingDeck.print()

	// fmt.Println(deck.toString())

	// deck.saveToFile("my_deck")

	myNewDeck := newDeckFromFile("my_deck")

	myNewDeck.suffle()

	myNewDeck.print()

}

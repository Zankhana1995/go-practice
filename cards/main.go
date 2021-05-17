package main

func main() {

	cards := newDeck()

	// iterate over slice and print
	// cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()

	// remainingDeck.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()

}

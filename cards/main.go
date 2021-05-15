package main

func main() {

	cards := newDeck()

	//iterate over slice and print
	//cards.print()

	hand, remainingDeck := deal(cards, 5)

	hand.print()

	remainingDeck.print()

}

package main

func main() {

	//slice
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterate over slice and print
	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}

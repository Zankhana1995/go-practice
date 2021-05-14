package main

import "fmt"

func main() {

	//slice
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//iterate over slice and print

	for i, card := range cards {
		fmt.Println(i, card)
	}
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}

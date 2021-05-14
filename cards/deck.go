package main

import "fmt"

// create a new type of 'deck'

// which is a slice of string
//(the new deck type extends (borrow) the behaviour of string)
type deck []string

// create and return cards, so add return type card
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// the argument after func keyword is called Receiver
// Any variable of type deck method will now gets access to print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}

func (d deck) toString() string {
	//convert decktype into slice of string
	//[]string(d)

	// join slice of string into a single string with comma seprated with strings package

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	// the last param is for permission
	// 0666 is anyone can read or write file
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {
	//return byteslice and error object
	// if nothing goes wrong while reading, the 'err' will be nil
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		// Option 1: log the error and return a call to newDeck()
		// Option 2: log the error and ebtirely quit the program
		fmt.Println("Error: ", err)

		// nonzero value indicates that something went wrong
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

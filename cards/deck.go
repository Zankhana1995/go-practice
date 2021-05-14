package main

import "fmt"

// create a new type of 'deck'

// which is a slice of string
//(the new deck type extends (borrow) the behaviour of string)
type deck []string

// the argument after func keyword is called Receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

package main

import "fmt"

//Create a new type of 'deck
//which is a slice of strings

type deck []string

//though not precisely, the above line can be understood as creating a new type that will extend functionality on top of slices (this is an OO understanding of what is going on)

func newDeck() deck {
	//this will not be a receiver function because when this is called we will not already be working with a deck
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + "of" + suit)
		}
	}

	// "_" is used above because initially the values were "i" and "j", but we never used them
	// "_" says 'I know there is a variable here, but I really don't care what its value is'

	return cards
}
func (d deck) print() {
	/*
		^^^any variable of type "deck" now has access to "print" method
		(d deck) is referred to as a 'receiver' on a function
			the first part of the receiver "d" is going to be the actual copy of the deck we're working with as it is available in the function
				 so in cards.print(), d == cards
			it serves a similar function to the 'this' keyword in javascript classes
			by convention, it is not the word 'this', but rather a short abbrev that matches the type of the receiver
	*/
	for i, card := range d {
		fmt.Println(i, card)
	}
}

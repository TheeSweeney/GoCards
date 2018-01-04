package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 permissions mean anyone can read/write this file
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) //ReadFile returns two values, the contents of a file and the error message, so, we assign the return values to two separate vars

	//if there is no error^^^, it will be assigned the value nil (similar to null)
	if err != nil {
		//Option 1 - log error and return a call to newDeck() - if we fail, let's at least return some deck to work with
		//Option 2 - log error and quit program - let's use this one
		fmt.Println("Error: ", err)
		os.Exit(1)
		//if ^^^ is passed anything but 0, it will exit and close the program
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//^^^long form method alternatively can be:
	// card := "Ace of Spades"

	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	/*
		^^^relying on Go compiler to infer that a var with the val "Ave of Spades" will be a string
		only use ':=' when assigning a value to a NEW variable
		after var is initialized it could be reassigned as such
			car = "Five of Diamonds"

		Variables can be initialized and later assigned
			var deckSize int
			deckSize = 53
		Variables can be initialized outside of a function, but cannot be assigned a value.
	*/

	/*
		var - tells go we're creating a variable; card - name of the variable; string - tells compiler that only type string will ever be assigned to this var; Ace of Spades - assign this value to var card

		Go is a statically typed language
			dynamically typed langauge (JS, Ruby, Python) don't care about the type of a variable
			statically typed languages (C++, Java, Go) DO care and will create an error if a variable is assigned a value of a type different to what the variable is initially set as
				main basic data types
					bool
					string
					int - number
					float64 - number with decimal
	*/

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	//go required us to tell it what data type a function will return ^^^
	return "Five of Diamonds"
}

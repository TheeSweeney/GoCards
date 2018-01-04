package main

func main() {
	// cards := []string{newCard(), "Ace of Diamonds"}
	// since I've created the deck type, I can ammend the above line to what is below
	//cards := deck{newCard(), "Ace of Diamonds"}
	//and now that I wrote a newDeck() function, I can again rewrite it as:
	cards := newDeck()
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

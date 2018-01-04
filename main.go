package main

func main() {
	// cards := []string{newCard(), "Ace of Diamonds"}
	// since I've created the deck type, I can ammend the above line to what is below
	//cards := deck{newCard(), "Ace of Diamonds"}
	//and now that I wrote a newDeck() function, I can again rewrite it as:
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// //deal returns 2 seperate values
	// //the first will be assigned to var "hand", the second to var "remainingDeck"

	// hand.print()
	// remainingDeck.print()
	//both 'hand' and 'remainingDeck' are of type 'deck' and so have access to the print function
	// greeting := "Hi there!"
	// fmt.Println(greeting)
	// >"Hi there!"
	// byteGreet := []byte(greeting)
	// fmt.Println(byteGreet)
	// >[72 105 32 116 104 101 114 101 33]

	cards.saveToFile("my_cards")

}

func newCard() string {
	return "Five of Diamonds"
}

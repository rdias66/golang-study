package main

import "fmt"

func main() {
	clubs := SuitBuilder("Clubs")
	diamonds := SuitBuilder("Diamonds")
	hearts := SuitBuilder("Hearts")
	spades := SuitBuilder("Spades")
	suits := []suit{clubs, diamonds, hearts, spades}
	deck := DeckBuilder(suits)

	DeckPrinter(deck)

	deckSize := len(deck)
	fmt.Println(deckSize)
}

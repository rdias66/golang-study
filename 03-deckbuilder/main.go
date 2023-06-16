package main

import "fmt"

func main() {
	clubs := SuitBuilder("Clubs")
	diamonds := SuitBuilder("Diamonds")
	hearts := SuitBuilder("Hearts")
	spades := SuitBuilder("Spades")
	var suitChecker int = 0
	if checkSuitSize(clubs)
		suitChecker++
	if checkSuitSize(diamonds)
		suitChecker++
	if checkSuitSize(hearts)
		suitChecker++
	if checkSuitSize(spades)
		suitChecker++

	var deck Deck
	
	if (suitChecker == 4) {
		suits := []suit{clubs, diamonds, hearts, spades}
		deck = DeckBuilder(suits)
	}
	else {
		fmt.Println("Suits size are not valid")
	}
	
	if checkDeckSize(deck)
		DeckPrinter(deck)
	
	if !checkDeckSize(deck)
		fmt.Println("Deck size invalid")
}

func checkSuitSize(s suit) bool{
	if(len(s) == 13){
		return true	
	}
	return false
}

func checkDeckSize(d deck) bool{
	if(len(d) == 52){
		return true	
	}
	return false
}

package main

import "fmt"

func main() {
	clubs := SuitBuilder("Clubs")
	diamonds := SuitBuilder("Diamonds")
	hearts := SuitBuilder("Hearts")
	spades := SuitBuilder("Spades")
	var suitChecker int 
	if checkSuitSize(clubs)
		suitChecker++
	if checkSuitSize(diamonds)
		suitChecker++
	if checkSuitSize(hearts)
		suitChecker++
	if checkSuitSize(spades)
		suitChecker++

	var deck Deck

	err := (suitChecker != 4)
	
	if (err) {
		fmt.Println("Suits size are not valid")
	}
	else {
		suits := []suit{clubs, diamonds, hearts, spades}
		deck = DeckBuilder(suits)
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

package main

import "fmt"

type deck []card

func DeckBuilder(suits []suit) deck {
	newDeck := deck{}
	for _, s := range suits {
		for _, c := range s {
			newDeck = append(newDeck, c)
		}
	}
	return newDeck
}

func DeckPrinter(deck deck) {
	for _, d := range deck {
		fmt.Println(d)
	}
}

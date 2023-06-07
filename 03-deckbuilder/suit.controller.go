package main

import "fmt"

type card string

type suit []card

func SuitBuilder(suitName string) suit {
	suitName = " " + suitName
	newSuit := suit{}
	for i := 1; i <= 13; i++ {
		if i == 1 {
			var newCard card = card("Ace of" + suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 11 {
			var newCard card = card("Jack of" + suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 12 {
			var newCard card = card("Queen of" + suitName)
			newSuit = append(newSuit, newCard)
		} else if i == 13 {
			var newCard card = card("King of" + suitName)
			newSuit = append(newSuit, newCard)
		} else {
			var newCard card = card(fmt.Sprint(i) + suitName)
			newSuit = append(newSuit, newCard)
		}
	}
	return newSuit
}

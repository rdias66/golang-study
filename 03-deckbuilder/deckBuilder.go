package main

type deck []card

func DeckBuilder(suits []suit) deck {
 newDeck := deck{}
  for _, s := range suits {
    for _, c := range s {
      newDeck = apped(newDeck, c)
    }
  }
}

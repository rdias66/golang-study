package main 

import (
    "fmt"
)

func main() {
  clubs := SuitBuilder("of Clubs")
  diamonds := SuitBuilder("of Diamonds")
  hearts := SuitBuilder("of Hearts")
  spades := SuitBuilder("of Spades")
  suits := []suit{clubs, diamonds, hearts, spades}
  deck := DeckBuilder(suits)
    
    //to be deck printer
    
    for _, d := range deck {
        fmt.Println(d)
    }
}

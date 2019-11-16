package main

import "fmt"

func main() {

	cards := newDeck()

	// fmt.Println(newDeckFromFile("deck-of-cards"))

	cards.shuffle()

	fmt.Println(cards)
}

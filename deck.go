package main

type deck []string

func newDeck() deck {

	var cards []string

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suite)
		}
	}

	return cards
}

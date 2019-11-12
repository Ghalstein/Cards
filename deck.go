package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

// creates a new deck of cards
func newDeck() deck {

	var cards []string

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suite)
		}
	}

	return deck(cards)
}

// reads the csv from a txt file and makes a deck
func newDeckFromFile(filename string) deck {

	byteslice, err := ioutil.ReadFile(filename)
	if err != nil {

		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteslice), ",") // will be a csv file of cards

	return deck(s)
}

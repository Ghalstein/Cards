package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

// shuffles the deck of cards
// takes the deck working and romizes the roder
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	// tests for the proper length of the deck
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// tests for the correct first index of the deck
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
}

package main

import "fmt"

func main() {
	fmt.Println(newDeck())

	fmt.Println(newDeckFromFile("deck-of-cards"))
}

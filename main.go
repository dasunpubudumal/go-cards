package main

import (
	"realitix.com/cards/evenodd"
)

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")

	readDeck := newDeckFromFile("cards.txt")

	readDeck.shuffleDeck()

	hand, _ := deal(readDeck, 5)

	hand.print()

	evenodd.EvenOrOdd([]int{1, 2, 3, 4, 5})

}

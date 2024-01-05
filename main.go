package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.txt")

	readDeck := newDeckFromFile("cards.txt")

	readDeck.shuffleDeck()

	hand, _ := deal(readDeck, 5)

	hand.print()
}

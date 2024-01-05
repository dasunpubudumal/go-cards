package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

// Doesn't require a receiver since it's not operating on anything
// It's just generating one deck
func newDeck() deck {
	cardDeck := deck{}
	for _, cardType := range []string{"Hearts", "Spades", "Clubs", "Diamonds"} {
		for _, num := range []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Ace", "General", "Queen", "King"} {
			cardDeck = append(cardDeck, fmt.Sprintf("%s of %s", num, cardType))
		}
	}
	return cardDeck
}

// d is a receiver of type deck
// receiver is usually a one or two letter abbreviation
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Create a new shuffled deck
// without updating the original deck
func (d deck) shuffleDeck() {
	rand.Shuffle(len(d),
		func(i, j int) { d[i], d[j] = d[j], d[i] })
}

// splits the card deck from 0 -> handSize and handSize -> len(deck)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Internal representation of the deck is a string of cards
// concatenated with "," as the delimiter.
func (d deck) saveToFile(fileName string) {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func newDeckFromFile(fileName string) deck {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data[:]), ",")
}

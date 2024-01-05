package main

import (
	"log"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	newDeck := newDeck()
	expected := 52
	if expected != len(newDeck) {
		t.Errorf("Length check failed. Found: %v, expected: %v", len(newDeck), expected)
	}
}

func TestShuffleDeck(t *testing.T) {
	newDeck := newDeck()
	expected := 52
	newDeck.shuffleDeck()
	if expected != len(newDeck) {
		t.Errorf("Length check failed. Found: %v, expected: %v", len(newDeck), expected)
	}
}

func TestDeal(t *testing.T) {
	newDeck := newDeck()
	hand, d := deal(newDeck, 5)
	if len(d) != (52 - 5) {
		t.Errorf("Length check failed for remaining hand. Expected: %v, Received: %v", 47, len(d))
	}
	if len(hand) != 5 {
		t.Errorf("Length check failed for hand. Expected: %v, remaining: %v", 5, len(d))
	}
}

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {
	fileName := "_deckfile"
	newDeck := newDeck()
	newDeck.saveToFile(fileName)
	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != 52 {
		t.Errorf("Length check failed for remaining hand. Expected: %v, Received: %v", 52, len(loadedDeck))
	}
	t.Cleanup(func() {
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal("Error in removing file.")
		}
	})
}

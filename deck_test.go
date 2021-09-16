package main

import "testing"
 

func TestNewDeck(t *testing.T){

	d:= newDeck()

	if len(d) != 16 {
		t.Errorf("New Deck should have 16 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect to get Ace of Spades , but got %v", d[0])
	}
}
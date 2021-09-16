package main

import "fmt"


type deck []string

func (d deck) print(){

	for i,cards := range d {
		fmt.Println(i, cards)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuit := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValue := []string{"Ace","Jack", "Queen", "King"}
	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}
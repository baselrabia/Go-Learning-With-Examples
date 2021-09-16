package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


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


func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	 oneString := strings.Join(d, ",")
	 return oneString
}

func toByte(s string) []byte {
	return []byte(s)
}


func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, toByte(d.toString()), 0666)
 
}

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

func newDeckFromFile(fileName string) deck {
	bs,err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:" , err)
		os.Exit(1)
	}
 
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano()) //seeder by time in nano
	r := rand.New(source) //generator

	for i := range d {
		newPosition := r.Intn(len(d)-1) //random by nano time 
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
package main

// import "fmt"

 
 

func main() {

//   cards := newDeck()

//   hand ,remaingCards := deal(cards,10)

//   hand.print()
//   remaingCards.print()

	// oneString := cards.toString()
	// oneString = string(toByte(oneString))
	// fmt.Println(oneString)

	// cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")

	cards.shuffle()
	cards.print()

}

package main

 
 

func main() {

  cards := newDeck()

  hand ,remaingCards := deal(cards,10)

  hand.print()
  remaingCards.print()


}

package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	//fromFile := newDeckFromFile("my_cards")
	//fromFile.print()

	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
}

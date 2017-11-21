package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	fromFile := newDeckFromFile("my_cards")
	fromFile.print()
}

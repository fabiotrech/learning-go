package main

func main() {
	cards := newDeckFromFile("my-deck.txt")
	cards.print()
}

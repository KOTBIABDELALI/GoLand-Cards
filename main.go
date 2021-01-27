package main

func main() {
	cards := deck{"newCard()", newCard()}

	cards = append(cards, "voila")
	cards.print()
}
func newCard() string {
	return "Ace of Spades"
}

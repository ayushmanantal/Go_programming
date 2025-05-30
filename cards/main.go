package main

func main() {
	// var card string = "Ace of Spades"
	// colon must only be used in initialisation of the variable in Go
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

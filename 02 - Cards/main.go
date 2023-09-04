package main


func main() {

	paragraph()

	shuffledDeck := getStandartShuffledDeck()
	shuffledDeck.print()

	paragraph()

	hand, shuffledDeck := shuffledDeck.getStdGameHand()

	shuffledDeck.print()

	paragraph()

    hand.showHand()
}

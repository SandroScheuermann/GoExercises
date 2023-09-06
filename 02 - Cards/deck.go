package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type deck []card

func getStandartShuffledDeck() deck {

	var stdDeck deck = deck(getStdCards())

	stdDeck.shuffleDeck()

	return stdDeck
}

func (d deck) deal(ammount int) ([]card, deck) {

	var cards []card

	fmt.Println("Qtd cartas no deck antes :", len(d))

	cards = append(d[:ammount])

	d = append(d[ammount:])

	fmt.Println("Qtd cartas no deck depois :", len(d))

	return cards, d
}

func (d deck) shuffleDeck() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) exportDeckToJson() error {

	jsonDeck, error := json.Marshal(d)

	byteSlice := []byte(jsonDeck)

	if error != nil {
		return error
	}

	return os.WriteFile("ExportedDecks/deckJson.json", byteSlice, 0666)
}

func readDeckJson() deck {

	byteSlice, error := os.ReadFile("ExportedDecks/deckJson.json")

	var decodedDeck deck

	error = json.Unmarshal(byteSlice, &decodedDeck)

	if error != nil {
		fmt.Println(error.Error())
	}

	return decodedDeck
}

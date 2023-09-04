package main

import ("fmt"; "math/rand")

type deck []card

func getStandartShuffledDeck() deck{

    var stdDeck deck = deck(getStdCards())

    stdDeck.shuffleDeck()

    return stdDeck
}

func (d deck) deal(ammount int) ([]card, deck){

    var cards []card

    cards = append(d[:ammount] )

    d = append(d[:ammount],d[ammount*2:]...)

    return cards,d 
}

func(d deck) shuffleDeck(){
    rand.Shuffle(len(d), func(i, j int){
        d[i], d[j] = d[j], d[i]
    })
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

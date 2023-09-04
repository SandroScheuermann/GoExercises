package main

import (
    "fmt"
)

type hand []card

func (d deck) getStdGameHand() (hand, deck) {

    var h hand

    h, d = d.deal(5)

    return h, d
}

func (h hand) showHand() {

    for i, card := range h {
        fmt.Println(i+1, "Carta da mão " + card)
    }

}

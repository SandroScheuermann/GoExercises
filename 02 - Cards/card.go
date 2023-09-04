package main

import (
	"math/rand"
)

type card string

var suits = []string{"Copas", "Ouros", "Paus", "Espadas"}
var values = []string{"Ás", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Dama", "Rei"}

func getStdCards() []card {

	var stdCards []card

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			stdCards = append(stdCards, card(values[j]+ " de " +suits[i]))
		}
	}

	return stdCards
}

func getRandomIndex(a []string) int {
	return rand.Intn(len(a))
}

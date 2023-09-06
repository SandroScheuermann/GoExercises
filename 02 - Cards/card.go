package main

import (
	"math/rand"
)

type card string

var suits = []string{"Copas", "Ouros", "Paus", "Espadas"}
var values = []string{"Ás", "Dois", "Três", "Quatro", "Cinco", "Seis", "Sete", "Oito", "Nove", "Dez", "Valete", "Dama", "Rei"}

func getStdCards() []card {

	var stdCards []card
    
    for _, suit := range suits{
        for _, value := range values{
            stdCards = append(stdCards, card(value + " de " + suit))
        }
    }

	return stdCards
}

func getRandomIndex(a []string) int {
	return rand.Intn(len(a))
}

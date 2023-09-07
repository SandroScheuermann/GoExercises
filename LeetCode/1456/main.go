package main

import (
	"fmt"
)

func main() {
	fmt.Print(maxVowels("aeiou", 2))
}

var vogais = map[rune]bool{

    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
}

func ehVogal(letra rune) bool{
    return vogais[letra]
}

func maxVowels(s string, k int) int {

    contadorMax, contadorJanela := 0, 0

    stringInicial := s[:k] 

    for _, letra := range stringInicial{
        if ehVogal(letra){
            contadorJanela++
        }
    }

    contadorMax = contadorJanela

    for i := k; i < len(s) ; i++{

        if ehVogal(rune(s[i-k])){
            contadorJanela--
        }

        if ehVogal(rune(s[i])){
            contadorJanela++ 
        }

        if(contadorJanela > contadorMax){
            contadorMax = contadorJanela
        }
    }

	return contadorMax 
}


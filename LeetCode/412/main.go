package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Tuple struct{
    n int
    w string
}


func main() {
	fmt.Printf("%v\n", strings.Join(fizzBuzz(1000), ","))
}

func fizzBuzz(n int) []string {

	m := [2]Tuple{{3, "Fizz"}, {5, "Buzz"}} 
    
	var answer []string

	for i := 1; i <= n; i++ {
		var currentWord string

		for _, mT := range m {
			if i%mT.n == 0 {
				currentWord += mT.w
			}
		}

		if currentWord == "" {
			currentWord = strconv.Itoa(i)
		}

		answer = append(answer, currentWord)
	}

	return answer
}

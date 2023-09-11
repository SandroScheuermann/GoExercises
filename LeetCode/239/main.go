package main

import "fmt"

type queue struct {
	number int
	indice int
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	fmt.Printf("\nesperado : [7,4] obtido : %d \n", maxSlidingWindow([]int{7, 2, 4}, 2))
}

func maxSlidingWindow(nums []int, k int) []int {

	var maxWindows []int
	var windowPositiveQueue queue
	windowPositiveQueue.number = MinInt

	for i := 0; i < len(nums); i++ {

		if nums[i] > windowPositiveQueue.number || k == 1{ // Verifica se o número atual iterado é maior que o  da janela positiva
			windowPositiveQueue.number = nums[i] // se for, salva ele na janela positiva
			windowPositiveQueue.indice = i + 1
		}
		if ((windowPositiveQueue.indice)%k == 0 && i > k) || (windowPositiveQueue.indice-1 == 0 && i >= k) { // maior numero guardado não pertence a janela atual

			// Devo pegar o indice do maior numero guardado + 1, comparar com o numero iterado atual, se for maior, dar o append
			// e colocar ele no maior numero guardado, se não, adicionar o numero iterado atual no append, e no maior numero guardado.
			// dar um continue para ele nao cair no fluxo normal
			if windowPositiveQueue.indice < len(nums) && nums[windowPositiveQueue.indice] > nums[i] && k > 1 {
				maxWindows = append(maxWindows, nums[windowPositiveQueue.indice])

				windowPositiveQueue.number = nums[windowPositiveQueue.indice]
				windowPositiveQueue.indice++
			} else {

				maxWindows = append(maxWindows, nums[i])

				windowPositiveQueue.number = nums[i]
				windowPositiveQueue.indice = i + 1
			}

			continue
		}

		// Se o número pertencer a janela, verificar se o indice i+1 é igual ou maior a k (tamanho da janela), se for, dar o append
		// para tratar a primeira janela K, e os appends normais caso o numero nao cair nos filtros de cima.

		if i+1 >= k {
			maxWindows = append(maxWindows, windowPositiveQueue.number)
		}

	}
	return maxWindows
}

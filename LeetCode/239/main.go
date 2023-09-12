package main

import (
	"fmt"
)

type queue []int

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	fmt.Printf("\nesperado : [3,3,2,5] obtido : %d \n", maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
    }

func maxSlidingWindow(nums []int, k int) []int {

	var maxWindows []int
	var indexQueue queue

	for i := 0; i < len(nums); i++ {
 
		for len(indexQueue) > 0 && nums[i] >= nums[indexQueue[len(indexQueue)-1]] {
			indexQueue = indexQueue.popLast()
		}
        
        indexQueue = append(indexQueue, i)

		if (i - k + 1) > indexQueue[0] {
			indexQueue = indexQueue.popFirst()
		}
        
		if i+1 >= k {
			maxWindows = append(maxWindows, nums[indexQueue[0]])
		}
	}

	return maxWindows
}

func (queueBuffer queue) popLast() queue {
    
	return queueBuffer[:len(queueBuffer)-1]
}

func (queueBuffer queue) popFirst() queue {

	return queueBuffer[1:]
}

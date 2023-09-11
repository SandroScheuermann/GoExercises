package main

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {

	obtido := maxSlidingWindow([]int{7, 2, 4}, 2)
	esperado := []int{7, 4}

	if !reflect.DeepEqual(obtido, esperado) {

		t.Errorf("\nObtido :%d \n Esperado : %d", obtido, esperado)
	}

}

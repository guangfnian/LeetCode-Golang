package p0033

import "testing"

func TestOK(t *testing.T) {
	inputsA := [][]int{{8, 0, 1, 2, 3, 4}}
	inputsB := []int{3}
	answers := []int{4}
	for i := range inputsA {
		if ret := search(inputsA[i], inputsB[i]); ret != answers[i] {
			//t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

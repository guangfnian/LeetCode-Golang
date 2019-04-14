package p0016

import "testing"

func TestOK(t *testing.T) {
	inputsA := [][]int{{-1, 2, 1, -4}}
	inputsB := []int{1}
	answers := []int{2}
	for i := range inputsA {
		if ret := threeSumClosest(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], " ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

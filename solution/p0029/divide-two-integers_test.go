package p0029

import "testing"

func TestOK(t *testing.T) {
	inputsA := []int{}
	inputsB := []int{}
	answers := []int{}
	for i := range inputsA {
		if ret := divide(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], " ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

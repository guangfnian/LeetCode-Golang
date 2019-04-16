package p0027

import "testing"

func TestOK(t *testing.T) {
	inputsA := [][]int{{3, 2, 2, 3}, {0, 1, 2, 2, 3, 0, 4, 2}}
	inputsB := []int{3, 2}
	answers := []int{2, 5}
	for i := range inputsA {
		if ret := removeElement(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], " ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

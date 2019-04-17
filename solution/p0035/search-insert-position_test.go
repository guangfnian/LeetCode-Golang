package p0035

import "testing"

func TestOK(t *testing.T) {
	inputsA := [][]int{{1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}, {1, 3, 5, 6}}
	inputsB := []int{5, 2, 7, 0}
	answers := []int{2, 1, 4, 0}
	for i := range inputsA {
		if ret := searchInsert(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], " + ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

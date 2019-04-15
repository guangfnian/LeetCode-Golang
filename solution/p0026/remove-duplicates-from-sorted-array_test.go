package p0026

import "testing"

func TestOK(t *testing.T) {
	inputs := [][]int{{1, 1, 2}, {0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}
	answers := []int{2, 5}
	for i := range inputs {
		if ret := removeDuplicates(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

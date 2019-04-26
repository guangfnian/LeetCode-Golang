package p0080

import "testing"

func TestOK(t *testing.T) {
	inputs := [][]int{{1, 1, 1, 2, 2, 3}, {0, 0, 1, 1, 1, 1, 2, 3, 3}}
	answers := []int{5, 7}
	for i := range inputs {
		if ret := removeDuplicates(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

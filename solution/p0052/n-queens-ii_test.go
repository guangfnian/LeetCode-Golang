package p0052

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{4}
	answers := []int{2}
	for i := range inputs {
		if ret := totalNQueens(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

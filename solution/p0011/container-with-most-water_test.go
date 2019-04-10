package p0011

import "testing"

func TestOK(t *testing.T) {
	inputs := [][]int{{1, 8, 6, 2, 5, 4, 8, 3, 7}}
	answers := []int{49}
	for i := range inputs {
		if ret := maxArea(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

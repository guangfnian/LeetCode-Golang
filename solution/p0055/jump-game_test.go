package p0055

import "testing"

func TestOK(t *testing.T) {
	inputs := [][]int{{2, 3, 1, 1, 4}, {3, 2, 1, 0, 4}}
	answers := []bool{true, false}
	for i := range inputs {
		if ret := canJump(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

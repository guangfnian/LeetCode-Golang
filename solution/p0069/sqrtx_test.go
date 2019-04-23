package p0069

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{4, 8}
	answers := []int{2, 2}
	for i := range inputs {
		if ret := mySqrt(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

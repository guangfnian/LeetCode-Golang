package p0007

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{123, -123, 120}
	answers := []int{321, -321, 21}
	for i := range inputs {
		if ret := reverse(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

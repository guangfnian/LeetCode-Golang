package p0009

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{121, -121, 10}
	answers := []bool{true, false, false}
	for i := range inputs {
		if ret := isPalindrome(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

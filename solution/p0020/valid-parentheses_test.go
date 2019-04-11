package p0020

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}
	answers := []bool{true, true, false, false, true}
	for i := range inputs {
		if ret := isValid(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

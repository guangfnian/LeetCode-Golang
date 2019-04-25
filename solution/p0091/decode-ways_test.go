package p0091

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"12", "226"}
	answers := []int{2, 3}
	for i := range inputs {
		if ret := numDecodings(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

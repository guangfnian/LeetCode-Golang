package p0014

import "testing"

func TestOK(t *testing.T) {
	inputs := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}}
	answers := []string{"fl", ""}
	for i := range inputs {
		if ret := longestCommonPrefix(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

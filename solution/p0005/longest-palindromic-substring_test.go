package p0005

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"cbbd", "a", "aaa", "abbab", "cabbad"}
	answers := []string{"bb", "a", "aaa", "abba", "abba"}

	for i := range inputs {
		if ret := longestPalindrome(inputs[i]); ret != answers[i] {
			t.Fatal("input: ", inputs[i], " want: ", answers[i], " got: ", ret)
		}
	}
}

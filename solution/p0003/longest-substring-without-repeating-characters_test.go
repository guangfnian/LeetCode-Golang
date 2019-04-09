package p0003

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"abcabcbb", "bbbbb", "pwwkew"}
	answers := []int{3, 1, 3}
	for i := range inputs {
		if ret := lengthOfLongestSubstring(inputs[i]); ret != answers[i] {
			t.Fatal(inputs[i], " wanted: ", answers[i], " got: ", ret)
		}
	}
}

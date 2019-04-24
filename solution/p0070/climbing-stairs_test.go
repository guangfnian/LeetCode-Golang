package p0070

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{2, 3}
	answers := []int{2, 3}
	for i := range inputs {
		if ret := climbStairs(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

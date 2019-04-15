package p0013

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	answers := []int{3, 4, 9, 58, 1994}
	for i := range inputs {
		if ret := romanToInt(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

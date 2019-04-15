package p0012

import "testing"

func TestOK(t *testing.T) {
	inputs := []int{3, 4, 9, 58, 1994}
	answers := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for i := range inputs {
		if ret := intToRoman(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

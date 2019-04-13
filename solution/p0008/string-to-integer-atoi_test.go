package p0008

import "testing"

func TestOK(t *testing.T) {
	inputs := []string{"42", "   -42", "4193 with words", "words and 987", "-91283472332"}
	answers := []int{42, -42, 4193, 0, -2147483648}
	for i := range inputs {
		if ret := myAtoi(inputs[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

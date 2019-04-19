package p0017

import (
	"reflect"
	"testing"
)

func TestOK(t *testing.T) {
	inputs := []string{"23"}
	answers := [][]string{{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}}
	for i := range inputs {
		if ret := letterCombinations(inputs[i]); !reflect.DeepEqual(ret, answers[i]) {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

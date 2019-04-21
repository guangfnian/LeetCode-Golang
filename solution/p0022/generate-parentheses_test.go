package p0022

import (
	"reflect"
	"testing"
)

func equal(a []string, b []string) bool {
	mpa := make(map[string]int)
	mpb := make(map[string]int)
	for _, i := range a {
		mpa[i]++
	}
	for _, i := range b {
		mpb[i]++
	}
	return reflect.DeepEqual(mpa, mpb)
}

func TestOK(t *testing.T) {
	inputs := []int{3}
	answers := [][]string{
		{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}}
	for i := range inputs {
		if ret := generateParenthesis(inputs[i]); !equal(ret, answers[i]) {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

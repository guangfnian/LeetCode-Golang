package p0051

import (
	"reflect"
	"testing"
)

//func euqal(x, y [][]string) bool

func TestOK(t *testing.T) {
	inputs := []int{4}
	answers := [][][]string{
		{
			{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
			{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
	}
	for i := range inputs {
		if ret := solveNQueens(inputs[i]); !reflect.DeepEqual(ret, answers[i]) {
			t.Fatal("inputs: ", inputs[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

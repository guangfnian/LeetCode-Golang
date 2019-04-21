package p0074

import "testing"

func TestOK(t *testing.T) {
	inputsA := [][][]int{
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
	}
	inputsB := []int{3, 13}
	answers := []bool{true, false}
	for i := range inputsA {
		if ret := searchMatrix(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], "+", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

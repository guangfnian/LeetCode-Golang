package p0034

import (
	"reflect"
	"testing"
)

func TestOK(t *testing.T) {
	inputsA := [][]int{{5, 7, 7, 8, 8, 10}, {5, 7, 7, 8, 8, 10}}
	inputsB := []int{8, 6}
	answers := [][]int{{3, 4}, {-1, -1}}
	for i := range inputsA {
		if ret := searchRange(inputsA[i], inputsB[i]); !reflect.DeepEqual(ret, answers[i]) {
			t.Fatal("inputs: ", inputsA[i], " + ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

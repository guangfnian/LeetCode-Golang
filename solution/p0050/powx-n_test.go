package p0050

import "testing"

func TestOK(t *testing.T) {
	inputsA := []float64{2, 2}
	inputsB := []int{10, -2}
	answers := []float64{1024, 0.25}
	for i := range inputsA {
		if ret := myPow(inputsA[i], inputsB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputsA[i], " , ", inputsB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

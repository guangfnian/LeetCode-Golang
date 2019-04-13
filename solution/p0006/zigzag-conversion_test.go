package p0006

import "testing"

func TestOK(t *testing.T) {
	inputA := []string{"PAYPALISHIRING", "PAYPALISHIRING"}
	inputB := []int{3, 4}
	answers := []string{"PAHNAPLSIIGYIR", "PINALSIGYAHRPI"}
	for i := range inputA {
		if ret := convert(inputA[i], inputB[i]); ret != answers[i] {
			t.Fatal("inputs: ", inputA[i], " ", inputB[i], " wants: ", answers[i], " got: ", ret)
		}
	}
}

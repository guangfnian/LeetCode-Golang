package p0002

import "testing"

func genNumber(arr []int) *ListNode {
	length := len(arr)
	if length == 0 {
		return nil
	}
	head := &ListNode{arr[0], nil}
	cur := head
	for i := 1; i < length; i++ {
		cur.Next = &ListNode{arr[i], nil}
		cur = cur.Next
	}
	return head
}

func equal(x, y *ListNode) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return false
	}
	if x.Val != y.Val {
		return false
	}
	return equal(x.Next, y.Next)
}

func TestFunc(t *testing.T) {
	inputA := [][]int{{2, 4, 3}, {0}, {9, 9, 9, 9}}
	inputB := [][]int{{5, 6, 4}, {0}, {1}}
	answers := [][]int{{7, 0, 8}, {0}, {0, 0, 0, 0, 1}}
	length := len(inputA)
	for i := 0; i < length; i++ {
		argA := genNumber(inputA[i])
		argB := genNumber(inputB[i])
		ans := genNumber(answers[i])
		res := addTwoNumbers(argA, argB)
		if !equal(ans, res) {
			t.Fatal(inputA[i], inputB[i])
		}
	}
}

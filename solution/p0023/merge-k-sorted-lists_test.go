package p0023

import "testing"

func TestHp_Len(t *testing.T) {
	n1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	n2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	n3 := &ListNode{2, &ListNode{6, nil}}
	arr := []*ListNode{n1, n2, n3}
	mergeKLists(arr)
}

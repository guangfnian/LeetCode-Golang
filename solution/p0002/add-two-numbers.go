package p0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	head := &ListNode{sum % 10, nil}
	add := sum / 10
	cur := head
	l1 = l1.Next
	l2 = l2.Next
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val + add
		add = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 == nil {
		l1 = l2
	}
	for l1 != nil {
		sum = l1.Val + add
		add = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
		l1 = l1.Next
	}
	if add != 0 {
		cur.Next = &ListNode{1, nil}
	}
	return head
}

package p0148

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(x, y *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	for x != nil && y != nil {
		if x.Val < y.Val {
			cur.Next = x
			x = x.Next
		} else {
			cur.Next = y
			y = y.Next
		}
		cur = cur.Next
	}
	for x != nil {
		cur.Next = x
		x = x.Next
		cur = cur.Next
	}
	for y != nil {
		cur.Next = y
		y = y.Next
		cur = cur.Next
	}
	return ret.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fast = slow
	slow = slow.Next
	fast.Next = nil
	head = sortList(head)
	slow = sortList(slow)
	return merge(head, slow)
}

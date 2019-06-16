package p0086

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	greater := head
	for greater != nil && greater.Val < x {
		greater = greater.Next
	}
	less := head
	for less != nil && less.Val >= x {
		less = less.Next
	}
	if greater == nil {
		return less
	}
	if less == nil {
		return greater
	}
	lHead := less
	gHead := greater
	var cur *ListNode
	if greater.Val == head.Val {
		cur = less.Next
		for greater.Next.Val != less.Val {
			greater = greater.Next
		}
	} else {
		cur = greater.Next
		for greater.Val != less.Next.Val {
			less = less.Next
		}
	}
	for cur != nil {
		if cur.Val >= x {
			greater.Next = cur
			greater = cur
		} else {
			less.Next = cur
			less = cur
		}
		cur = cur.Next
	}
	less.Next = gHead
	greater.Next = nil
	return lHead
}

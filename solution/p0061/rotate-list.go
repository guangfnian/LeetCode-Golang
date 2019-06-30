package p0061

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	cur := head
	tail := head
	for cur != nil {
		tail = cur
		cur = cur.Next
		length++
	}
	if length == 0 {
		return head
	}
	k = length - (k % length)
	if k == length {
		return head
	}
	cur = head
	pre := head
	count := 0
	for count < k {
		pre = cur
		cur = cur.Next
		count++
	}
	pre.Next = nil
	tail.Next = head
	return cur
}

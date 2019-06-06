package p0024

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	nxt := head.Next
	head.Next = swapPairs(nxt.Next)
	nxt.Next = head
	return nxt
}

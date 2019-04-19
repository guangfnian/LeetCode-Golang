package p0019

type ListNode struct {
	Val  int
	Next *ListNode
}

func dfs(cur *ListNode, n int) (*ListNode, int) {
	if cur == nil {
		return nil, 1
	}
	next, cnt := dfs(cur.Next, n)
	cur.Next = next
	if cnt == n {
		return next, cnt + 1
	}
	return cur, cnt + 1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head, _ = dfs(head, n)
	return head
}

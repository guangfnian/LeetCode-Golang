package p0023

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Hp []*ListNode

func (h Hp) Len() int {
	return len(h)
}

func (h Hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h *Hp) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Hp) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return r
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := make(Hp, 0)
	for i := range lists {
		if lists[i] != nil {
			h = append(h, lists[i])
		}
	}
	head := &ListNode{}
	cur := head
	heap.Init(&h)
	for h.Len() > 0 {
		p := heap.Pop(&h).(*ListNode)
		cur.Next = p
		cur = cur.Next
		if p.Next != nil {
			heap.Push(&h, p.Next)
		}
	}
	return head.Next
}

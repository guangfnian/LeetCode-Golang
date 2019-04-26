package p0103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	T *TreeNode
	L int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	que := make([]*Node, 1)
	que[0] = &Node{root, 0}
	for len(que) > 0 {
		node := que[0].T
		level := que[0].L
		que = que[1:]
		if len(ret) == level {
			ret = append(ret, make([]int, 0))
		}
		if level&1 == 0 {
			ret[level] = append(ret[level], node.Val)
		} else {
			ret[level] = append([]int{node.Val}, ret[level]...)
		}
		if node.Left != nil {
			que = append(que, &Node{node.Left, level + 1})
		}
		if node.Right != nil {
			que = append(que, &Node{node.Right, level + 1})
		}
	}
	return ret
}

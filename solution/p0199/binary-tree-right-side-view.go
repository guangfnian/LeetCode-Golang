package p0199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	T *TreeNode
	L int
}

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	que := make([]*Node, 1)
	que[0] = &Node{root, 0}
	for len(que) != 0 {
		node := que[0].T
		level := que[0].L
		que = que[1:]
		if level >= len(ret) {
			ret = append(ret, node.Val)
		} else {
			ret[level] = node.Val
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

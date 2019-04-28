package p0662

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	T *TreeNode
	L int
	P int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	l := make([]int, 0)
	r := make([]int, 0)
	que := make([]*Node, 1)
	que[0] = &Node{root, 0, 0}
	for len(que) > 0 {
		node := que[0].T
		level := que[0].L
		pos := que[0].P
		que = que[1:]
		if level == len(l) {
			l = append(l, pos)
			r = append(r, pos)
		}
		if l[level] > pos {
			l[level] = pos
		}
		if r[level] < pos {
			r[level] = pos
		}
		ret = max(ret, r[level]-l[level])
		if node.Left != nil {
			que = append(que, &Node{node.Left, level + 1, pos + pos + 1})
		}
		if node.Right != nil {
			que = append(que, &Node{node.Right, level + 1, pos + pos + 2})
		}
	}
	return ret + 1
}

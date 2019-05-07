package p0101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(x, y *TreeNode, flag *bool) {
	if *flag == false {
		return
	}
	if x == nil {
		if y != nil {
			*flag = false
		}
		return
	}
	if y == nil {
		*flag = false
		return
	}
	if x.Val != y.Val {
		*flag = false
		return
	}
	dfs(x.Left, y.Right, flag)
	dfs(x.Right, y.Left, flag)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	dfs(root.Left, root.Right, &flag)
	return flag
}

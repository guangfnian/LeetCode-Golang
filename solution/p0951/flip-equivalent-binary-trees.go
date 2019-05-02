package p0951

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil {
		if root2 != nil {
			return false
		}
		return true
	}
	if root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	if (dfs(root1.Left, root2.Left) && dfs(root1.Right, root2.Right)) || (dfs(root1.Left, root2.Right) && dfs(root1.Right, root2.Left)) {
		return true
	}
	return false
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	return dfs(root1, root2)
}

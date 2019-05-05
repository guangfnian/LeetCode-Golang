package p0114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var pre *TreeNode = nil
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		dfs(root.Left)
		root.Left = nil
		root.Right = pre
		pre = root
	}
	dfs(root)
}

package p0129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, cur int, ans *int) {
	if root == nil {
		return
	}
	cur = root.Val + cur
	if root.Left == nil && root.Right == nil {
		*ans += cur
		return
	}
	cur *= 10
	dfs(root.Left, cur, ans)
	dfs(root.Right, cur, ans)
}

func sumNumbers(root *TreeNode) int {
	ret := 0
	dfs(root, 0, &ret)
	return ret
}

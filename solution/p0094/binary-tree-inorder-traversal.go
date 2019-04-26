package p0094

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	ret = append(ret, root.Val)
	dfs(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	ret = make([]int, 0)
	dfs(root)
	return ret
}

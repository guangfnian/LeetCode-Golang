package p0979

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dfs(root *TreeNode, ans *int) (int, int) {
	// return total node, coin
	if root == nil {
		return 0, 0
	}
	n1, c1 := dfs(root.Left, ans)
	n2, c2 := dfs(root.Right, ans)
	*ans += abs(n1-c1) + abs(n2-c2)
	return n1 + n2 + 1, c1 + c2 + root.Val
}

func distributeCoins(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans
}

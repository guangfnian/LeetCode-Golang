package p0513

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	depth := -1
	pos := 0
	ret := 0
	var dfs func(rt *TreeNode, d, p int)
	dfs = func(rt *TreeNode, d, p int) {
		if rt == nil {
			return
		}
		if d > depth {
			depth = d
			pos = p
			ret = rt.Val
		} else if d == depth && p < pos {
			pos = p
			ret = rt.Val
		}
		dfs(rt.Left, d+1, (p<<1)+1)
		dfs(rt.Right, d+1, (p<<1)+2)
	}
	dfs(root, 0, 0)
	return ret
}

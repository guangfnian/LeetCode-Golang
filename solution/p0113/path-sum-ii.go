package p0113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, cur, target int, rec []int, ret *[][]int) {
	if root == nil {
		return
	}
	cur += root.Val
	rec = append(rec, root.Val)
	if root.Left == nil && root.Right == nil {
		if cur == target {
			tmp := make([]int, len(rec))
			copy(tmp, rec)
			*ret = append(*ret, tmp)
		}
		return
	}
	dfs(root.Left, cur, target, rec, ret)
	dfs(root.Right, cur, target, rec, ret)
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	rec := make([]int, 0)
	dfs(root, 0, sum, rec, &ret)
	return ret
}

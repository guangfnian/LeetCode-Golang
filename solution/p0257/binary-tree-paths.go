package p0257

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	val := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{val}
	}
	var left []string
	var right []string
	if root.Right != nil {
		right = binaryTreePaths(root.Right)
		for i := range right {
			right[i] = val + "->" + right[i]
		}
	}
	if root.Left != nil {
		left = binaryTreePaths(root.Left)
		for i := range left {
			left[i] = val + "->" + left[i]
		}
	}
	var ret []string
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret
}

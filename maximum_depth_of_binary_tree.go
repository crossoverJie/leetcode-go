package leetcode_go

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	var left, right int
	left = maxDepth(root.Left)
	right = maxDepth(root.Right)

	return max(left, right) + 1
}

func max(left, right int) int {
	if left > right {
		return left
	} else {
		return right
	}
}

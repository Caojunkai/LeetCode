package leetcode

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
const INT_MIN = 1

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0, maxPathSum(root.Left))
	right := max(0, maxPathSum(root.Right))

	println(left, right)

	return 0
}

func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}

// @lc code=end

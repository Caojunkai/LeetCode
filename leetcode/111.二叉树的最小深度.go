/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */
package leetcode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLeft := minDepth(root.Left)
	minRight := minDepth(root.Right)

	if minLeft == 0 && minRight != 0 {
		return minRight + 1
	}

	if minLeft != 0 && minRight == 0 {
		return minLeft + 1
	}

	return min111(minLeft, minRight) + 1
}

func min111(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

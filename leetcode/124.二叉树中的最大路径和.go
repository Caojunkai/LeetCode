/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

var maxPath int

func maxPathSum(root *TreeNode) int {
	maxPath = -1 << 31
	getMaxPath(root)
	return maxPath
}

func getMaxPath(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := max(getMaxPath(root.Left), 0)
	right := max(getMaxPath(root.Right), 0)

	maxPath = max(left+right+root.Val, maxPath)

	return root.Val + max(left, right)
}

func max(data ...int) int {
	if len(data) == 0 {
		return -1 << 31
	}
	resp := data[0]
	for _, v := range data {
		if resp < v {
			resp = v
		}
	}

	return resp
}

// @lc code=end

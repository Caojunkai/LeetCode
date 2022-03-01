/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
var maxTravel int

func diameterOfBinaryTree(root *TreeNode) int {
	maxTravel = 0
	maxTravelLength(root)
	return maxTravel
}

func maxTravelLength(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxTravelLength(root.Left)
	right := maxTravelLength(root.Right)

	maxTravel = max543(maxTravel, left+right)

	return max543(left, right) + 1

}

func max543(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

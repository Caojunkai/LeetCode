/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
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

func levelOrderBottom(root *TreeNode) [][]int {
	var levels [][]int
	var findLevel func(root *TreeNode, level int)

	findLevel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(levels) == level {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], root.Val)

		findLevel(root.Left, level+1)
		findLevel(root.Right, level+1)
	}

	findLevel(root, 0)

	length := len(levels)
	end := length / 2

	for i := 0; i < end; i++ {
		levels[i], levels[length-i-1] = levels[length-i-1], levels[i]
	}

	return levels
}

// @lc code=end

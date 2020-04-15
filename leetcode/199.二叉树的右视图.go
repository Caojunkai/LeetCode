/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
var rsp []int

func rightSideView(root *TreeNode) []int {
	rsp = []int{}
	dfs(root, 1)
	return rsp
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level > len(rsp) {
		rsp = append(rsp, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}

// @lc code=end

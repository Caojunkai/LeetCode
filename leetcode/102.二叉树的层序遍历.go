/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	data := make(map[int][]int)
	travel(root, 0, data)

	resp := make([][]int, len(data))
	for k, v := range data {
		resp[k] = v
	}
	return resp
}

func travel(root *TreeNode, level int, data map[int][]int) {
	if root == nil {
		return
	}
	data[level] = append(data[level], root.Val)
	travel(root.Left, level+1, data)
	travel(root.Right, level+1, data)
	return
}

// @lc code=end

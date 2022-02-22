/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	levelMap := make(map[int]int)
	levelMaxValue(root, 0, levelMap)
	resp := make([]int, len(levelMap))
	for k, v := range levelMap {
		resp[k] = v
	}
	return resp
}

func levelMaxValue(root *TreeNode, level int, data map[int]int) {
	if root == nil {
		return
	}

	if temp, ok := data[level]; !ok || root.Val > temp {
		data[level] = root.Val
	}

	levelMaxValue(root.Left, level+1, data)
	levelMaxValue(root.Right, level+1, data)
}

// @lc code=end

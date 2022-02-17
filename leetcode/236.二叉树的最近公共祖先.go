/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */
package leetcode

// @lc code=start

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p.Val == root.Val || q.Val == root.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左子树为空, 说明其中两个值都不在左子树, 它们一定都在右子树
	if left == nil {
		return right
	}

	// 如果右子树为空, 说明其中两个值都不在右子树, 它们一定都左子树
	if right == nil {
		return left
	}

	// 如果左右子树都不为nil, 说明两个值分别在树的两侧, 根节点就为最近的公共祖先
	return root
}

// @lc code=end

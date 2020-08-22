/**

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

package lcof

import "github.com/fitzix/leetcode/classic"

func reConstructBinaryTree(pre []int, vin []int) *classic.TreeNode {
	if len(pre) == 0 {
		return nil
	}

	root := &classic.TreeNode{
		Val: pre[0],
	}

	if len(pre) == 1 {
		return root
	}

	mid := findMiddle(vin, pre[0])

	root.Left = reConstructBinaryTree(pre[1:mid+1], vin[:mid])
	root.Right = reConstructBinaryTree(pre[mid+1:], vin[mid+1:])
	return root
}

func findMiddle(vin []int, target int) int {
	for _, v := range vin {
		if v == target {
			return v
		}
	}
	return 0
}

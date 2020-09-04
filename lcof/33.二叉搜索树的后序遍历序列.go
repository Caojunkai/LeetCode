// https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/

/*

输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。



参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3
示例 1：

输入: [1,6,3,2,5]
输出: false
示例 2：

输入: [1,3,2,6,5]
输出: true

*/

package lcof

func verifyPostorder(postorder []int) bool {
	return verfify(postorder, 0, len(postorder)-1)
}

func verfify(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}

	middle := start
	for postorder[middle] < postorder[end] {
		middle++
	}
	right := middle
	for postorder[right] > postorder[end] {
		right++
	}
	return right == end && verfify(postorder, start, middle-1) && verfify(postorder, middle, end-1)
}

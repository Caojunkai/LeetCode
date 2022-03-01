/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */
package leetcode

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	head := new(ListNode)

	l, l1, l2 := head, list1, list2

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			l.Next = l2
			l2 = l2.Next
		} else {
			l.Next = l1
			l1 = l1.Next
		}
		l = l.Next
	}

	if l1 == nil {
		l.Next = l2
	}

	if l2 == nil {
		l.Next = l1
	}

	return head.Next
}

// @lc code=end

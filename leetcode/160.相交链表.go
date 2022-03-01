/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	l1, l2 := headA, headB
	l1Done, l2Done := false, false

	for l1 != l2 {
		l1 = l1.Next
		if l1 == nil && !l1Done {
			l1Done = true
			l1 = headB
		}

		l2 = l2.Next
		if l2 == nil && !l2Done {
			l2Done = true
			l2 = headA
		}

	}
	return l1
}

// @lc code=end

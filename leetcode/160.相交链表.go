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

	for l1 != l2 {
		if next := l1.Next; next != nil {
			l1 = next
		} else {
			l1 = headB
		}

		if next := l2.Next; next != nil {
			l2 = next
		} else {
			l2 = headA
		}
	}

	return l1
}

// @lc code=end

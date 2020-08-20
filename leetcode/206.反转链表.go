/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
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
func reverseList(head *ListNode) *ListNode {
	// var pre, next *ListNode
	// for head != nil {
	// 	next = head.Next
	// 	head.Next = pre
	// 	pre = head
	// 	head = next
	// }
	// return pre
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// @lc code=end

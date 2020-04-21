/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
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
func middleNode(head *ListNode) *ListNode {
	p, q := head, head
	for q != nil {
		if q.Next == nil {
			break
		}
		p = p.Next
		q = q.Next.Next
	}
	return p
}

// @lc code=end

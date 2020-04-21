/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := new(ListNode)
	l.Next = head
	p, q := l, l

	for i := 0; i <= n; i++ {
		q = q.Next
	}

	for q != nil {
		q = q.Next
		p = p.Next
	}

	p.Next = p.Next.Next
	return l.Next
}

// @lc code=end

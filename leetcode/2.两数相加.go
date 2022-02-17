/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rsp *ListNode
	var needNextCount bool
	p, q, cur := l1, l2, rsp

	for p != nil || q != nil {
		var l1Val, l2Val int

		if p != nil {
			l1Val = p.Val
			p = p.Next
		}

		if q != nil {
			l2Val = q.Val
			q = q.Next
		}

		total := l1Val + l2Val

		if needNextCount {
			total += 1
		}

		node := &ListNode{
			Val: total % 10,
		}

		if total >= 10 {
			needNextCount = true
		} else {
			needNextCount = false
		}

		if cur == nil {
			cur = node
			rsp = cur
			continue
		}

		cur.Next = node
		cur = node
	}

	if needNextCount {
		cur.Next = &ListNode{
			Val: 1,
		}
	}

	return rsp
}

// @lc code=end

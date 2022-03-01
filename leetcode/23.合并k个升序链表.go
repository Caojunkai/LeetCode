/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
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
import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	h := new(Heap)
	for _, v := range lists {
		heap.Push(h, v)
	}
	head := new(ListNode)
	p := head
	for h.Len() > 0 {
		temp := heap.Pop(h)
		p.Next = temp.(*ListNode)
		p = p.Next
		if p != nil {
			heap.Push(h, p)
		}
	}
	return head.Next
}

type Heap []*ListNode

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	resp := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return resp
}

// @lc code=end

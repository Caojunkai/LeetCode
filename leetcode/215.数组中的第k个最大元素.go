/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */
package leetcode

import (
	"container/heap"
)

// @lc code=start
func findKthLargest(nums []int, k int) int {
	return nums[find(nums, 0, len(nums)-1, k-1)]
}

func find(nums []int, left, right, k int) int {
	mid := part(nums, left, right)
	if mid > k {
		return find(nums, left, mid-1, k)
	}
	if mid < k {
		return find(nums, mid+1, right, k)
	}
	return mid
}

func part(nums []int, left, right int) int {
	if left >= right {
		return left
	}

	i, j := left+1, right
	for i <= j {
		if nums[i] < nums[left] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	i--
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

// @lc code=end

func FindKthLargestWithHeap(nums []int, k int) int {
	h := new(myHeap)
	for i := 0; i < len(nums); i++ {
		if h.Len() < k {
			heap.Push(h, nums[i])
			continue
		}
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}

	return (*h)[0]
}

type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m myHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *myHeap) Push(v interface{}) {
	*m = append(*m, v.(int))
}

func (m *myHeap) Pop() interface{} {
	res := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return res
}

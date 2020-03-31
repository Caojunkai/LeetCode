/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */
package leetcode

// nolint
// @lc code=start
func sortArray(nums []int) []int {
	QuickSort(nums, 0, len(nums)-1)
	return nums
}

func QuickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left+1, right
	for i <= j {
		if nums[i] > nums[left] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	i--
	nums[left], nums[i] = nums[i], nums[left]
	QuickSort(nums, left, i-1)
	QuickSort(nums, i+1, right)
}

// @lc code=end

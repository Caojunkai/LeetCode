/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package leetcode

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	ret := make([]int, len(nums1))
	var i, j, k int
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			ret[k] = nums1[i]
			i++
		} else {
			ret[k] = nums2[j]
			j++
		}
		k++
	}
	if i == m {
		for _, v := range nums2[j:m] {
			ret[k] = v
			k++
		}
	}

	if j == n {
		for _, v := range nums1[i:n] {
			ret[k] = v
			k++
		}
	}
	nums1 = ret
}

// @lc code=end

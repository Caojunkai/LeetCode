/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */
package leetcode

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {

	resp := make([]int, 0, m+n)

	i, j := 0, 0

	for i < m && j < n {
		if nums1[i] >= nums2[j] {
			resp = append(resp, nums2[j])
			j++
		} else {
			resp = append(resp, nums1[i])
			i++
		}
	}

	if i >= m {
		resp = append(resp, nums2[j:n]...)
	}

	if j >= n {
		resp = append(resp, nums1[i:m]...)
	}

	copy(nums1, resp)
}

// @lc code=end

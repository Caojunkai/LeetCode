/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] 山脉数组的峰顶索引
 */
package leetcode

// @lc code=start
func peakIndexInMountainArray(arr []int) int {

	start, end := 0, len(arr)-1

	for start <= end {
		mid := (start + end) / 2

		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid] >= arr[mid-1] {
			start = mid
		} else {
			end = mid
		}
	}

	return -1
}

// @lc code=end
